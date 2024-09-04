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
	url := "https://openapi.koreainvestment.com:9443//uapi/domestic-stock/v1/quotations/inquire-price"

	token, err := GenerateToken(appKey, appSecret)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("토큰", token)

	var rtn map[string]interface{}

	header := map[string]string{
		"content-type":  "application/json; charset=utf-8",
		"authorization": "Bearer " + token.accessToken,
		"appkey":        appKey,
		"appsecret":     appSecret,
		"tr_id":         "FHKST01010100",
	}

	body := map[string]string{
		"FID_COND_MRKT_DIV_CODE": "J",
		"FID_INPUT_ISCD":         "005930",
	}

	err = sendRequest(url, http.MethodGet, header, body, &rtn)
	if err != nil {

	}

	fmt.Println(rtn)
}
