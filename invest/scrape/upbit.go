package scrape

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (s Scraper) upbitApi(sym string) (float64, error) {

	url := s.t.ApiBaseUrl("upbit")
	if url == "" {
		return 0, errors.New("URL 미존재")
	}
	url = fmt.Sprintf(url, sym)

	var rtn []map[string]string
	err := sendRequest(url, http.MethodGet, nil, nil, &rtn)
	if err != nil {
		return 0, err
	}

	cp, err := strconv.ParseFloat(rtn[0]["trade_price"], 64)
	if err != nil {
		return 0, err
	}

	return cp, nil
}
