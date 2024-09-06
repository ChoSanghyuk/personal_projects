package scrape

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expired     string `json:"access_token_token_expired"`
}

func (s *Scraper) KisToken() (string, error) {

	if s.kis.accessToken != "" && strings.Compare(s.kis.tokenExpired, time.Now().Format("2006-01-02 15:04:05")) == 1 {
		return s.kis.accessToken, nil
	}

	url := "https://openapi.koreainvestment.com:9443/oauth2/tokenP"

	var token TokenResponse
	err := sendRequest(url, http.MethodPost, nil, map[string]string{
		"grant_type": "client_credentials",
		"appkey":     s.kis.appKey,
		"appsecret":  s.kis.appSecret,
	}, &token)
	if err != nil {
		return "", err
	}

	s.kis.accessToken = token.AccessToken
	s.kis.tokenExpired = token.Expired

	return token.AccessToken, nil
}

type CurrentPriceResponse struct {
	CurrentPrice string `json:"stck_prpr"`
}

func (s *Scraper) kisDomesticStockCurrentPrice(code string) (float64, error) {

	// url := "https://openapi.koreainvestment.com:9443/uapi/domestic-stock/v1/quotations/inquire-price?fid_cond_mrkt_div_code=J&fid_input_iscd=" + code
	url := s.t.ApiBaseUrl("upbit")
	if url == "" {
		return 0, errors.New("URL 미존재")
	}
	url = fmt.Sprintf(url, code)

	token, err := s.KisToken()
	if err != nil {
		return 0, err
	}

	var rtn CurrentPriceResponse

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer " + token,
		"appkey":        s.kis.appKey,
		"appsecret":     s.kis.appSecret,
		"tr_id":         "FHKST01010100",
	}

	err = sendRequest(url, http.MethodGet, header, nil, &rtn)
	if err != nil {
		return 0, err
	}

	cp, err := strconv.ParseFloat(rtn.CurrentPrice, 64)
	if err != nil {
		return 0, err
	}

	return cp, nil
}
