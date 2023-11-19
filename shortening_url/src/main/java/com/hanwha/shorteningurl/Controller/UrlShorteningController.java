package com.hanwha.shorteningurl.Controller;

import java.io.IOException;
import java.time.LocalDateTime;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.hanwha.shorteningurl.model.Url;
import com.hanwha.shorteningurl.model.UrlDto;
import com.hanwha.shorteningurl.model.UrlErrorResponseDto;
import com.hanwha.shorteningurl.model.UrlResponseDto;
import com.hanwha.shorteningurl.service.UrlService;

@RestController
public class UrlShorteningController {

	private final String BASEURL = "http://localhost:8080/";

	@Autowired
	private UrlService urlService;
	
	// client로 부터 생성 요청을 받아 처리합니다. 
	@PostMapping("/generate")
	public ResponseEntity<?> genereateShortLink(@RequestBody UrlDto urlDto, HttpServletRequest request){
			
		Url urlToReturn = urlService.generateShortLink(urlDto); // url with a shortlink created
		
		// Url 객체가 잘 생성되었다면, 해당 정보가 담긴 UrlResponseDto 객체를 생성하여 client에게 반환합니다. 
		if(urlToReturn != null) {
			UrlResponseDto urlResponseDto = new UrlResponseDto();
			urlResponseDto.setOriginalUrl(urlToReturn.getOriginalUrl());
			urlResponseDto.setExpirationDate(urlToReturn.getExpirationDate());
			urlResponseDto.setShortLink(BASEURL+urlToReturn.getShortLink());
			return new ResponseEntity<UrlResponseDto>(urlResponseDto, HttpStatus.OK);
		}
		
		// 만약 Url 객체가 생성되지 못 했다면 오류를 반환합니다.
		UrlErrorResponseDto urlErrorResponseDto = new UrlErrorResponseDto();
        urlErrorResponseDto.setStatus("400");
        urlErrorResponseDto.setError("잘못된 url 형식입니다. 다시 입력해주세요.");
        return new ResponseEntity<UrlErrorResponseDto>(urlErrorResponseDto,HttpStatus.OK);
	}
	
	// client가 짧은 url을 입력했을 때, 원본 url 페이지로 redirect 해줍니다.
	@GetMapping("/{shortLink}")
	public ResponseEntity<?> redirectToOriginalUrl(@PathVariable String shortLink, HttpServletResponse response) throws IOException{
		
		// url이 비어 있을 경우 오류 처리입니다.
		if(StringUtils.isEmpty(shortLink)) {	
			UrlErrorResponseDto urlErrorResponseDto = new UrlErrorResponseDto();
            urlErrorResponseDto.setError("잘못된 url 형식입니다. 다시 입력해주세요.");
            urlErrorResponseDto.setStatus("400");
            return new ResponseEntity<UrlErrorResponseDto>(urlErrorResponseDto,HttpStatus.OK);
		}
		
		Url urlToReturn = urlService.getEncodedUrl(shortLink);
		// 해당 짧은 url 가진 Url 객체가 없는 경우입니다. 처음부터 저장된 url이 아니거나 주기적인 만료 url 삭제 작업을 통해 삭제되었을 수 있습니다.
        if(urlToReturn == null){
            UrlErrorResponseDto urlErrorResponseDto = new UrlErrorResponseDto();
            urlErrorResponseDto.setError("url이 존재하지 않거나 만료되었습니다."); // url이 존재하지 않음
            urlErrorResponseDto.setStatus("404");
            return new ResponseEntity<UrlErrorResponseDto>(urlErrorResponseDto,HttpStatus.OK);
        }
		
        // db에는 남아있지만, 만료 기간이 이미 지난 경우입니다.
        if(urlToReturn.getExpirationDate().isBefore(LocalDateTime.now())) {
            urlService.deleteShortLink(urlToReturn);
            UrlErrorResponseDto urlErrorResponseDto = new UrlErrorResponseDto();
            urlErrorResponseDto.setError("url이 만료되었습니다. 새로운 url을 발급 받으십시오.");
            urlErrorResponseDto.setStatus("404");
            return new ResponseEntity<UrlErrorResponseDto>(urlErrorResponseDto,HttpStatus.OK);
        }
        // 위의 오류 사항들에 해당하지 않는다면, 원본 url 사이트로 redirect 해줍니다.
        response.sendRedirect(urlToReturn.getOriginalUrl());
        return null;
		
		
	}
	
	
}
