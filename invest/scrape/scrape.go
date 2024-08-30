package scrape

import (
	"encoding/json"
	"fmt"
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
