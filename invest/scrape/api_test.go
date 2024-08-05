package scrape

import (
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
	var info = config.NewConfigInfo()

	url := info.Bitcoin.API.Url
	id := info.Bitcoin.API.ID
	key := info.Bitcoin.API.ApiKey

	rtn, err := s.CallApi(url, map[string]string{
		"X-Naver-Client-Id":     id,
		"X-Naver-Client-Secret": key,
	})
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}
