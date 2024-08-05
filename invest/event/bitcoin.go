package event

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type BitcoinEventHandler struct {
	Scraper interface {
		Crawl(url string, cssPath string) (string, error)
	}
	Url        string
	Csspath    string
	UpperBound float64
	LowerBound float64
}

func (b BitcoinEventHandler) PullingEvent(c chan<- string) {

	for true {
		rtn, err := b.Scraper.Crawl(b.Url, b.Csspath)
		if err != nil {
			fmt.Println(err.Error())
			c <- fmt.Sprintf("크롤링 시 오류 발생. %s", err.Error())
			continue
		}

		bitPrice, err := strconv.ParseFloat(strings.ReplaceAll(rtn, ",", ""), 64)
		if err != nil {
			log.Print(err.Error())
			c <- fmt.Sprintf("크롤링 시 오류 발생. %s", err.Error())
			continue
		}

		if bitPrice >= b.UpperBound {
			c <- fmt.Sprintf("SELL BIT. UPPER BOUND : %f. CURRENT PRICE :%f", b.UpperBound, bitPrice)

		} else if bitPrice <= b.LowerBound {
			c <- fmt.Sprintf("BUY BIT. LOWER BOUND : %f. CURRENT PRICE :%f", b.LowerBound, bitPrice)
		} else {
			log.Printf("비트코인 현재 가격 %.3f", bitPrice)
		}
		time.Sleep(10 * time.Minute)
	}
}
