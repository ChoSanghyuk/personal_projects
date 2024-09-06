package scrape

import (
	"errors"
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
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
