package scrape

import (
	"net/http"
	"strings"
	"time"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expired     string `json:"access_token_token_expired"`
}

func (s *Scraper) KisToken() (string, error) {

	if s.KIS.accessToken != "" && strings.Compare(s.KIS.tokenExpired, time.Now().Format("2006-01-02 15:04:05")) == 1 {
		return s.KIS.accessToken, nil
	}

	url := "https://openapi.koreainvestment.com:9443/oauth2/tokenP"

	var token TokenResponse
	err := sendRequest(url, http.MethodPost, nil, map[string]string{
		"grant_type": "client_credentials",
		"appkey":     s.KIS.appKey,
		"appsecret":  s.KIS.appSecret,
	}, &token)
	if err != nil {
		return "", err
	}

	s.KIS.accessToken = token.AccessToken
	s.KIS.tokenExpired = token.Expired

	return token.AccessToken, nil
}

type CurrentPriceResponse struct {
	CurrentPrice string `json:"stck_prpr"`
}

func (s *Scraper) KisCurrentPrice(target string) (string, error) {
	url := "https://openapi.koreainvestment.com:9443/uapi/domestic-stock/v1/quotations/inquire-price?fid_cond_mrkt_div_code=J&fid_input_iscd=" + target

	token, err := s.KisToken()
	if err != nil {
		return "", err
	}

	var rtn CurrentPriceResponse

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer " + token,
		"appkey":        s.KIS.appKey,
		"appsecret":     s.KIS.appSecret,
		"tr_id":         "FHKST01010100",
	}

	err = sendRequest(url, http.MethodGet, header, nil, &rtn)
	if err != nil {
		return "", err
	}

	return rtn.CurrentPrice, nil
}
