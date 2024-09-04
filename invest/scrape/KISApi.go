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

func KISApi() {
	url := "https://openapi.koreainvestment.com:9443//uapi/domestic-stock/v1/quotations/inquire-price"

	_ = url

}
