package com.hanwha.shorteningurl.model;

//데이터 전송 객체(Data Transfer Object)로 client한테의 Url 수신에 사용될 객체입니다.
public class UrlDto
{
    private String url;
    private String expirationDate;  

    public UrlDto(String url, String expirationDate) {
        this.url = url;
        this.expirationDate = expirationDate;
    }

    public UrlDto() {
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getExpirationDate() {
        return expirationDate;
    }

    public void setExpirationDate(String expirationDate) {
        this.expirationDate = expirationDate;
    }

    @Override
    public String toString() {
        return "UrlDto{" +
                "url='" + url + '\'' +
                ", expirationDate='" + expirationDate + '\'' +
                '}';
    }
}
