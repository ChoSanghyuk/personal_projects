# 한화생명 테스트 과제 : URL Shortener 구현



## 해야할거

- expired 된 url 제거 작업 예약하기



- 변환 요청이 url이 맞는지 확인
- url 앞에 프로토콜 명시하지 않은 경우, 자동으로 추가 (http://)

- 만료기간 하루



## 프로젝트 구조

### shorteningurl

- src/main/java
  - com.hanwha.shorteningurl
    - `UrlshorternerApplication.java`
  - com.hanwha.shorteningurl.Controller
    - `UrlShorteningController.java`
  - com.hanwha.shorteningurl.model
    - `Url.java`
    - `UrlDto.java`
    - `UrlErrorResponseDto.java`
    - `UrlResponseDto.java`
  - com.hanwha.shorteningurl.repository
    - `UrlRepository.java`
  - com.hanwha.shorteningurl.service
    - `UrlService.java`
    - `UrlServiceImpl.java`
- src/main/resources
  - `application.properties`
- `pom.xml`



## 프로젝트 세팅 



### 사용 dependency

- spring-boot-starter-data-jpa
- spring-boot-starter-web
- h2
- guava
- commons-lang3
- spring-boot-devtools



### DB 세팅





## 기능 구현



### Model



### service



### repository



### Controller

