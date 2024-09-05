package scrape

import (
	"fmt"
	"net/http"
	"time"
)

type Token struct {
	accessToken string
	createdAt   time.Time
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expired     string `json:"access_token_token_expired"`
}

func GenerateToken(appKey string, appSecret string) (*Token, error) {

	url := "https://openapi.koreainvestment.com:9443/oauth2/tokenP"

	var rtn TokenResponse
	err := sendRequest(url, http.MethodPost, nil, map[string]string{
		"grant_type": "client_credentials",
		"appkey":     appKey,
		"appsecret":  appSecret,
	}, &rtn)
	if err != nil {
		return nil, err
	}

	fmt.Println(rtn)

	return &Token{
		accessToken: rtn.AccessToken,
		createdAt:   time.Now(),
	}, nil
}

func KISApi(appKey string, appSecret string) {
	url := "https://openapi.koreainvestment.com:9443/uapi/domestic-stock/v1/quotations/inquire-price-2"

	// token, err := GenerateToken(appKey, appSecret)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	var rtn map[string]interface{}

	header := map[string]string{
		"Content-Type":  "application/json",
		"authorization": "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6IjY1OTYwZWNmLWQ3ZDYtNGE3MC1hMmIwLTc1ZTZhY2Y4YWQ5OCIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNTU3OTA5MywiaWF0IjoxNzI1NDkyNjkzLCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.Vlla2RpNeBf40aXkos-9_MnKQhoi0oLGUnCPyyROd3iZorKnlIaOMqhekB9o3iECfcUcaVImxtI9tg5DnqRwrg",
		"appkey":        appKey,
		"appsecret":     appSecret,
		"tr_id":         "FHKST01010100",
	}

	url = url + "?fid_cond_mrkt_div_code=J&fid_input_iscd=005930"

	err := sendRequest(url, http.MethodGet, header, nil, rtn)
	if err != nil {
		fmt.Println("에러", err)
	}

	fmt.Println(rtn)
}
