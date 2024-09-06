package scrape

import (
	"errors"
	"fmt"
	"invest/model"
	"regexp"
	"strconv"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/gofiber/fiber/v2/log"
)

type Scraper struct {
	exchange struct {
		Rate float64
		Date time.Time
	}
	kis struct {
		appKey       string
		appSecret    string
		accessToken  string
		tokenExpired string
	}
	t transmitter
}

type transmitter interface {
	ApiBaseUrl(target string) string
	CrawlUrlCasspath(target string) (url string, cssPath string)
}

func NewScraper(t transmitter, options ...func(*Scraper)) *Scraper {
	s := &Scraper{
		t: t,
	}

	for _, opt := range options {
		opt(s)
	}
	return s
}

func WithKIS(appKey string, appSecret string) func(*Scraper) {

	return func(s *Scraper) {
		s.kis.appKey = appKey
		s.kis.appSecret = appSecret
	}
}

func WithToken(token string) func(*Scraper) {

	return func(s *Scraper) {
		s.kis.accessToken = token
		s.kis.tokenExpired = time.Now().Add(time.Duration(1) * time.Hour).Format("2006-01-02 15:04:05")
	}
}

func AlpacaCrypto(target string) (string, error) {

	client := marketdata.NewClient(marketdata.ClientOpts{})
	request := marketdata.GetCryptoBarsRequest{
		TimeFrame: marketdata.OneMin,
		Start:     time.Now().Add(time.Duration(-10) * time.Minute), // time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		End:       time.Now(),
	}

	bars, err := client.GetCryptoBars(target, request)
	if err != nil {
		return "", err
	}

	if len(bars) == 0 {
		return "", errors.New("빈 결과값 반환")
	}

	bar := bars[len(bars)-1]
	return fmt.Sprintf("%f", bar.Close), nil
}

/*
종목 이름만 보고 어디서 가져올 지 정할 수 있어야 함
종목별로 타입을 지정 => 어떤 base url을 사용할 지 결정

	어떤 base url일지는 CurrentPrice 내부에서 case 세분화

종목별로 심볼 등 base url에 들어갈 인자를 정할 수 있어야함

종목 이름 - 타입/심볼을 어디에 저장해 둘 것인가 => DB
*/
func (s *Scraper) CurrentPrice(category model.Category, code string) (float64, error) {

	switch category {
	case model.DomesticStock:
		return s.kisDomesticStockCurrentPrice(code)
	case model.DomesticCoin:
		return s.upbitApi(code)
	}

	return 0, errors.New("미분류된 종목")
}

func (s *Scraper) RealEstateStatus() (string, error) {
	url, cssPath := s.t.CrawlUrlCasspath("estate")
	return s.crawl(url, cssPath)
}

func (s *Scraper) ExchageRate() float64 {

	if s.exchange.Rate != 0 && s.exchange.Date.Format("20060102") == time.Now().Format("20060102") {
		return s.exchange.Rate
	}

	// Todo config화 시킬지 결정
	url := "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%ED%99%98%EC%9C%A8"
	cssPath := "#main_pack > section.sc_new.cs_nexchangerate > div:nth-child(1) > div.exchange_bx._exchange_rate_calculator > div > div.inner > div:nth-child(3) > div.num > div > span"

	rtn, err := s.crawl(url, cssPath)
	if err != nil {
		log.Error(err)
	}

	re := regexp.MustCompile(`[^\d.]`)
	exrate, err := strconv.ParseFloat(re.ReplaceAllString(rtn, ""), 64)
	if err != nil {
		return 0
	}

	return exrate // TODO 환율 크롤링
}
