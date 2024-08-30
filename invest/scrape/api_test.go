package scrape

import (
	"encoding/json"
	"fmt"
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoldApi(t *testing.T) {

	var s = Scraper{}
	var info = config.NewConfigInfo()

	url := info.Gold.API.Url
	key := info.Gold.API.ApiKey

	rtn, err := s.CallApi(url, map[string]string{"x-access-token": key})
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}

func TestBitcoinApi(t *testing.T) {

	var s = Scraper{}
	// var info = config.NewConfigInfo()

	url := "https://api.upbit.com/v1/candles/minutes/1?market=KRW-BTC&count=1" //info.Bitcoin.API.Url
	// id := info.Bitcoin.API.ID
	// key := info.Bitcoin.API.ApiKey

	rtn, err := s.CallApi(url, nil)
	if err != nil {
		t.Error(err)
	}

	var d map[string]any
	json.Unmarshal([]byte(rtn[1:len(rtn)-1]), &d)

	assert.NotEmpty(t, rtn)

	t.Logf(fmt.Sprintf("%f", d["trade_price"]))
}
