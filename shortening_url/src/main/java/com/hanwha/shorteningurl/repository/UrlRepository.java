package com.hanwha.shorteningurl.repository;

import java.time.LocalDateTime;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import com.hanwha.shorteningurl.model.Url;

// 서버와 DB간의 교류를 담당합니다. JpaRepository를 상속받음으로 기본적인 저장과 삭제를 정의없이 사용합니다.
@Repository
public interface UrlRepository extends JpaRepository<Url, Long> {

	// 짧은 url를 통해 원본 url 정보가 담긴 Url 객체를 가져오는 메소드입니다.
	public Url findByShortLink(String shortLink);
	
	// query문을 통해 만료시간이 지난 url을 주기적으로 DB에서 삭제합니다.
	@Modifying
	@Query(value="delete from Url where expiration_date <= :now", nativeQuery=true)
	public void deleteExpiredAll(@Param("now") LocalDateTime now);
}
