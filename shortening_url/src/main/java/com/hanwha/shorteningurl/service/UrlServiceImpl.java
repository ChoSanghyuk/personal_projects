package com.hanwha.shorteningurl.service;

import java.nio.charset.StandardCharsets;
import java.time.LocalDate;
import java.time.LocalDateTime;

import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Transactional;

import com.google.common.hash.Hashing;
import com.hanwha.shorteningurl.model.Url;
import com.hanwha.shorteningurl.model.UrlDto;
import com.hanwha.shorteningurl.repository.UrlRepository;

@Component
public class UrlServiceImpl implements UrlService{
	
	@Autowired
	private UrlRepository urlRepository;

	@Override
	public Url generateShortLink(UrlDto urlDto) {
		
		// client로 부터 온 데이터에 url 정보가 비어있다면 null을 반환합니다.
		if(StringUtils.isNotEmpty(urlDto.getUrl())) {
			
			// client로 받은 데이터가 url 데이터인지 확인합니다. 만약 프로토콜을 빼먹었다면 앞에 붙여서 저장합니다.
			if(urlDto.getUrl().startsWith("https://") || urlDto.getUrl().startsWith("http://")) {
				// pass
			} else if (urlDto.getUrl().startsWith("www.")) {
				urlDto.setUrl("https://"+urlDto.getUrl());
			} else {
				return null;
			}
			
			// urlDto에서 받은 정보를 사용하여 Url 객체를 생성 및 저장합니다.
			String encodedUrl = encodeUrl(urlDto.getUrl());
			// 변환된 url을 db에 저장하기 위한 작업
			Url urlToPersist = new Url();
			urlToPersist.setCreationDate(LocalDateTime.now());
            urlToPersist.setOriginalUrl(urlDto.getUrl());
            urlToPersist.setShortLink(encodedUrl);
            urlToPersist.setExpirationDate(getExpirationDate(urlDto.getExpirationDate(),urlToPersist.getCreationDate()));
            Url urlToReturn = saveShortLink(urlToPersist);
			
            if(urlToReturn != null) {
            	return urlToReturn;
            }
		}
		
		return null;
	}
	
	// 별도의 만료날짜가 주어진다면 해당 만료날짜를 반환하고, 그렇지 않다면 생성시간에서 하루를 추가하여 만료 시간을 설정합니다.
	private LocalDateTime getExpirationDate(String expirationDate, LocalDateTime creationDate) {
		
		if(StringUtils.isBlank(expirationDate)) {
			return creationDate.plusHours(24);
		}
		
		LocalDateTime expirationDateToReturn = LocalDateTime.parse(expirationDate);
		return expirationDateToReturn;
	}


	// 긴 url을 짧은 url로 바꾸는 역할을 수행합니다.
	// 이때 같은 url에 대한 변환 요청이 있어도 다른 짧아진 url을 반환하기 위해 현재 시각을 사용합니다.
	private String encodeUrl(String url) {
		
		String encodedUrl = "";
		LocalDateTime time = LocalDateTime.now();
		encodedUrl = Hashing.murmur3_32()			// 짧은 url로 변환하기 위해 해당 해시 알고리즘을 선택했습니다.
				.hashString(url.concat(time.toString()), StandardCharsets.UTF_8)	// url과 현재 시간을 더해 유니크한 변환 값을 만들어냅니다.
				.toString();
		
		return encodedUrl;
	}


	// url 객체를 db에 저장합니다.
	@Override
	public Url saveShortLink(Url url) {
		Url urlToReturn = urlRepository.save(url);
		return urlToReturn;
	}

	// 짧은 url을 통해 원본 url 정보가 담긴 Url 객체를 반환합니다.
	@Override
	public Url getEncodedUrl(String url) {
		Url urlToReturn = urlRepository.findByShortLink(url);
		
		return urlToReturn;
	}

	// 해당 짧은 url을 가진 Url 객체를 db에서 삭제합니다.
	@Override
	public void deleteShortLink(Url url) {
		urlRepository.delete(url);
		
	}

	// 만료된 url을 삭제하는 작업을 매일 새벽 2시에 예약해둡니다. 이 시간은 서버의 트래픽 양이 제일 적은 시간대로 설정할 수 있습니다.
	@Override
	@Transactional
	@Scheduled(cron = "0 0 2 * * *")
	public void deleteExpiredShortLink() {
		urlRepository.deleteExpiredAll(LocalDateTime.now());
		
	}

}
