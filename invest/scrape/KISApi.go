package scrape

import (
	"fmt"
	"time"
)

type Token struct {
	accessToken string
	createdAt   time.Time
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expired     string `json:"acess_token_token_expired"`
}

func GenerateToken(appKey string, appSecret string) (*Token, error) {

	url := "https://openapi.koreainvestment.com:9443/oauth2/tokenP"

	var rtn TokenResponse
	err := sendRequest(url, map[string]string{
		"grant_type": "client_credentials",
		"appkey":     appKey,
		"appsecret":  appSecret,
	}, nil, &rtn)
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
	url := "https://openapi.koreainvestment.com:9443//uapi/domestic-stock/v1/quotations/inquire-price"

	t, err := GenerateToken(appKey, appSecret)
	fmt.Println("토큰", t)

	var rtn map[string]interface{}

	header := map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + t.accessToken,
		"appkey":        appKey,
		"appsecret":     appSecret,
		"tr_id":         "FHKST01010100",
	}

	body := map[string]string{
		"FID_COND_MRKT_DIV_CODE": "J",
		"FID_INPUT_ISCD":         "005930",
	}

	err = sendRequest(url, header, body, &rtn)
	if err != nil {

	}

	fmt.Println(rtn)
}
