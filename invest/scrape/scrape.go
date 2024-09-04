package scrape

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

type Scraper struct {
	ScrapeOption func() (string, error)
}

func NewScraper(option func() (string, error)) *Scraper {

	return &Scraper{
		ScrapeOption: option,
	}
}

func (s *Scraper) Scrape() (string, error) {
	return s.ScrapeOption()
}

func BitcoinApi(url string) func() (string, error) {

	return func() (string, error) {
		rtn, err := callApi(url, nil)
		if err != nil {
			return "", err
		}
		var d map[string]any

		err = json.Unmarshal([]byte(rtn[1:len(rtn)-1]), &d)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%f", d["trade_price"]), nil
	}
}

func (s *Scraper) GetRealtimeExchageRate() float64 {
	return 1337.58 // TODO 환율 크롤링
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

func KISApi() {

}
