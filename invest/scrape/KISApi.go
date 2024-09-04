package scrape

import (
	"fmt"
	"time"
)

type Token struct {
	accessToken string
	createdAt   time.Time
}

func GenerateToken(appKey string, appSecret string) (*Token, error) {

	url := "https://openapi.koreainvestment.com:9443/oauth2/tokenP"

	rtn, err := callApi(url, map[string]string{
		"grant_type": "client_credentials",
		"appkey":     appKey,
		"appsecret":  appSecret,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(rtn)

	return &Token{
		accessToken: "",
		createdAt:   time.Now(),
	}, nil
}

func KISApi() {

}
