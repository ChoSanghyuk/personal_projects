package com.hanwha.shorteningurl.service;

import org.springframework.stereotype.Service;

import com.hanwha.shorteningurl.model.Url;
import com.hanwha.shorteningurl.model.UrlDto;

@Service
public interface UrlService
{
    public Url generateShortLink(UrlDto urlDto);	// 짧은 url을 생성합니다.
    public Url saveShortLink(Url url);				// 생성된 url을 저장시킵니다.
    public Url getEncodedUrl(String url);			// 짧은 url을 통해 원본 url 정보가 담긴 Url 객체를 반환합니다.
    public void deleteShortLink(Url url);			// 생성한 url을 제거합니다.
    public void deleteExpiredShortLink();			// 만료된 url 객체를 db에서 제거합니다.
    
}