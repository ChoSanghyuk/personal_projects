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

type StockPrice struct {
	cp float64
	ap float64
	hp float64
	lp float64
}

func (s *Scraper) kisDomesticStockPrice(code string) (StockPrice, error) {

	url := s.t.ApiBaseUrl("KIS")
	if url == "" {
		return StockPrice{}, errors.New("URL 미존재")
	}
	url = fmt.Sprintf(url, code)

	token, err := s.KisToken()
	if err != nil {
		return StockPrice{}, err
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
		return StockPrice{}, err
	}

	if rtn.RtCd != "0" {
		return StockPrice{}, errors.New("국내 주식현재가 시세 API 실패 코드 반환")
	}

	cp, err := strconv.ParseFloat(rtn.Output["stck_prpr"], 64)
	if err != nil {
		return StockPrice{}, err
	}

	ap, err := strconv.ParseFloat(rtn.Output["wghn_avrg_stck_prc"], 64) // 가중 평균 주식 가격
	if err != nil {
		return StockPrice{}, err
	}

	hp, err := strconv.ParseFloat(rtn.Output["w52_hgpr"], 64)
	if err != nil {
		return StockPrice{}, err
	}

	lp, err := strconv.ParseFloat(rtn.Output["w52_lwpr"], 64)
	if err != nil {
		return StockPrice{}, err
	}

	return StockPrice{
		cp: cp,
		ap: ap,
		hp: hp,
		lp: lp,
	}, nil
}

// 해외주식 종목/지수/환율기간별시세(일/주/월/년)[v1_해외주식-012]
/*
해당 API로 미국주식 조회 시, 다우30, 나스닥100, S&P500 종목만 조회 가능합니다.
더 많은 미국주식 종목 시세를 이용할 시에는, 해외주식기간별시세 API
*/
func (s *Scraper) kisNasdaqIndex() (float64, error) {

	today := time.Now().Format("20060102")
	url := fmt.Sprintf("https://openapi.koreainvestment.com:9443/uapi/overseas-price/v1/quotations/inquire-daily-chartprice?FID_COND_MRKT_DIV_CODE=N&FID_INPUT_ISCD=COMP&FID_INPUT_DATE_1=%s&FID_INPUT_DATE_2=%s&FID_PERIOD_DIV_CODE=D", today, today)

	token, err := s.KisToken()
	if err != nil {
		return 0, err
	}

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer " + token,
		"appkey":        s.kis.appKey,
		"appsecret":     s.kis.appSecret,
		"tr_id":         "FHKST03030100",
	}

	type NasdaqResp struct {
		Msg    string `json:"msg1"`
		MsgCd  string `json:"msg_cd"`
		RtCd   string `json:"rt_cd"`
		Output struct {
			PresentPrice string `json:"ovrs_nmix_prpr"`
		} `json:"output1"` // value가 string 타입으로 넘어오기에 바로 파싱 X
	}
	var rtn NasdaqResp //TempResp

	err = sendRequest(url, http.MethodGet, header, nil, &rtn)
	if err != nil {
		return 0, err
	}

	if rtn.RtCd != "0" {
		return 0, errors.New("나스닥 인덱스 API 조회 실패 코드 반환")
	}

	cp, err := strconv.ParseFloat(rtn.Output.PresentPrice, 64)
	if err != nil {
		return 0, err
	}

	return cp, nil
}
