package com.hanwha.shorteningurl.model;

import java.time.LocalDateTime;

// client한테 url 변환 요청을 왔을 시, 성공적으로 변환되었다면 반환할 객체입니다.
public class UrlResponseDto
{
    private String originalUrl;
    private String shortLink;
    private LocalDateTime expirationDate;

    public UrlResponseDto(String originalUrl, String shortLink, LocalDateTime expirationDate) {
        this.originalUrl = originalUrl;
        this.shortLink = shortLink;
        this.expirationDate = expirationDate;
    }

    public UrlResponseDto() {
    }

    public String getOriginalUrl() {
        return originalUrl;
    }

    public void setOriginalUrl(String originalUrl) {
        this.originalUrl = originalUrl;
    }

    public String getShortLink() {
        return shortLink;
    }

    public void setShortLink(String shortLink) {
        this.shortLink = shortLink;
    }

    public LocalDateTime getExpirationDate() {
        return expirationDate;
    }

    public void setExpirationDate(LocalDateTime expirationDate) {
        this.expirationDate = expirationDate;
    }

    @Override
    public String toString() {
        return "UrlResponseDto{" +
                "originalUrl='" + originalUrl + '\'' +
                ", shortLink='" + shortLink + '\'' +
                ", expirationDate=" + expirationDate +
                '}';
    }
}
