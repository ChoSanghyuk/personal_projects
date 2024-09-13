package scrape

import (
	"errors"
	"fmt"
	"log"
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

	log.Println(token.AccessToken)

	s.kis.accessToken = token.AccessToken
	s.kis.tokenExpired = token.Expired

	return token.AccessToken, nil
}

type KIsResp struct {
	Msg    string            `json:"msg1"`
	MsgCd  string            `json:"msg_cd"`
	Output map[string]string `json:"output"` // value가 string 타입으로 넘어오기에 바로 파싱 X
	RtCd   string            `json:"rt_cd"`
}

func (s *Scraper) kisDomesticStockPrice(code string) (float64, float64, float64, error) {

	url := s.t.ApiBaseUrl("KIS")
	if url == "" {
		return 0, 0, 0, errors.New("URL 미존재")
	}
	url = fmt.Sprintf(url, code)

	token, err := s.KisToken()
	if err != nil {
		return 0, 0, 0, err
	}

	var rtn KIsResp

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer " + token,
		"appkey":        s.kis.appKey,
		"appsecret":     s.kis.appSecret,
		"tr_id":         "FHKST01010100",
	}

	err = sendRequest(url, http.MethodGet, header, nil, &rtn)
	if err != nil {
		return 0, 0, 0, err
	}

	if rtn.RtCd != "0" {
		return 0, 0, 0, errors.New("국내 주식현재가 시세 API 실패 코드 반환")
	}

	cp, err := strconv.ParseFloat(rtn.Output["stck_prpr"], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	hp, err := strconv.ParseFloat(rtn.Output["w52_hgpr"], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	lp, err := strconv.ParseFloat(rtn.Output["w52_lwpr"], 64)
	if err != nil {
		return 0, 0, 0, err
	}

	return cp, hp, lp, nil
}

func (s *Scraper) kisForeingIndex() {

	url := "https://openapi.koreainvestment.com:9443/uapi/overseas-price/v1/quotations/inquire-daily-chartprice?FID_COND_MRKT_DIV_CODE=N&FID_INPUT_ISCD=COMP&FID_INPUT_DATE_1=20240911&FID_INPUT_DATE_2=20240913&FID_PERIOD_DIV_CODE=D"

	token, err := s.KisToken()
	if err != nil {
		return
	}

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer " + token,
		"appkey":        s.kis.appKey,
		"appsecret":     s.kis.appSecret,
		"tr_id":         "FHKST03030100",
	}

	// type TempResp struct {
	// 	Msg    string            `json:"msg1"`
	// 	MsgCd  string            `json:"msg_cd"`
	// 	Output map[string]string `json:"output1"` // value가 string 타입으로 넘어오기에 바로 파싱 X
	// 	RtCd   string            `json:"rt_cd"`
	// }
	var rtn map[string]any //TempResp

	err = sendRequest(url, http.MethodGet, header, nil, &rtn)
	if err != nil {
		return
	}

	fmt.Printf("%+v", rtn)
}
