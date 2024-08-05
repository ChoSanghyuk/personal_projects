package event

import (
	"fmt"
	"log"
	"time"
)

type RealEstateEventHandler struct {
	Scraper interface {
		Crawl(url string, cssPath string) (string, error)
	}
	Url     string
	Csspath string
}

func (r RealEstateEventHandler) PullingEvent(c chan<- string) {

	for true {
		rtn, err := r.Scraper.Crawl(r.Url, r.Csspath)
		if err != nil {
			c <- fmt.Sprintf("크롤링 시 오류 발생. %s", err.Error())
			continue
		}

		if rtn != "예정지구 지정" {
			c <- fmt.Sprintf("연신내 재개발 변동 사항 존재. 예정지구 지정 => %s", rtn)
		} else {
			log.Printf("연신내 변동 사항 없음. 현재 단계: %s", rtn)
		}
		time.Sleep(1 * time.Hour)
	}
}
