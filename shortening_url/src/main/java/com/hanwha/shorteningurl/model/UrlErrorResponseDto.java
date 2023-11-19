package com.hanwha.shorteningurl.model;

//client한테 url 변환 요청을 왔을 시, 실패했다면 반환할 오류 메시지 객체입니다.
public class UrlErrorResponseDto
{
    private String status;
    private String error;

    public UrlErrorResponseDto(String status, String error) {
        this.status = status;
        this.error = error;
    }

    public UrlErrorResponseDto() {
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }

    @Override
    public String toString() {
        return "UrlErrorResponseDto{" +
                "status='" + status + '\'' +
                ", error='" + error + '\'' +
                '}';
    }
}
