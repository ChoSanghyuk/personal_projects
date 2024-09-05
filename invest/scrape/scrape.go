package scrape

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/gofiber/fiber/v2/log"
)

type Scraper struct {
	Exchange struct {
		Rate float64
		Date time.Time
	}
	KIS struct {
		appKey       string
		appSecret    string
		accessToken  string
		tokenExpired string
	}
}

func NewScraper(options ...func(*Scraper)) *Scraper {
	s := &Scraper{}

	for _, opt := range options {
		opt(s)
	}
	return s
}

func WithKIS(appKey string, appSecret string) func(*Scraper) {

	return func(s *Scraper) {
		s.KIS.appKey = appKey
		s.KIS.appSecret = appSecret
	}
}

func WithToken(token string) func(*Scraper) {

	return func(s *Scraper) {
		s.KIS.accessToken = token
		s.KIS.tokenExpired = time.Now().Add(time.Duration(1) * time.Hour).Format("2006-01-02 15:04:05")
	}
}

func (s *Scraper) ExchageRate() float64 {

	if s.Exchange.Rate != 0 && s.Exchange.Date.Format("20060102") == time.Now().Format("20060102") {
		return s.Exchange.Rate
	}

	// Todo config화 시킬지 결정
	url := "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%ED%99%98%EC%9C%A8"
	cssPath := "#main_pack > section.sc_new.cs_nexchangerate > div:nth-child(1) > div.exchange_bx._exchange_rate_calculator > div > div.inner > div:nth-child(3) > div.num > div > span"

	rtn, err := s.Crawl(url, cssPath)
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
