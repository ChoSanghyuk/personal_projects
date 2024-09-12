package scrape

import (
	"invest/config"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoldApi(t *testing.T) {

	info, _ := config.NewConfig()

	url := info.Api["gold"].Url
	head := info.Api["gold"].Header

	var rtn map[string]interface{}
	err := sendRequest(url, http.MethodGet, head, nil, rtn)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}

func TestBitcoinApi(t *testing.T) {

	c, _ := config.NewConfig()
	s := NewScraper(c)

	rtn, err := s.upbitApi("KRW-BTC")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%f", rtn)
}

func TestAlpaca(t *testing.T) {
	cp, err := AlpacaCrypto("BTC/USD")
	if err != nil {
		t.Error(err)
	}
	t.Log(cp)
}

func TestGoldCrwal(t *testing.T) {

	s := Scraper{}
	conf, _ := config.NewConfig()

	url := conf.Crawl["gold"].Url
	cssPath := conf.Crawl["gold"].CssPath

	rtn, err := s.crawl(url, cssPath)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}

func TestBitcoinCrwal(t *testing.T) {

	s := Scraper{}
	conf, _ := config.NewConfig()

	t.Run("Crwal", func(t *testing.T) {
		url := conf.Crawl["bitcoin"].Url
		cssPath := conf.Crawl["bitcoin"].CssPath

		rtn, err := s.crawl(url, cssPath)
		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, rtn)

		t.Log(rtn)
	})

}

func TestEstateCrwal(t *testing.T) {

	s := Scraper{}
	conf, _ := config.NewConfig()

	t.Run("Crwal", func(t *testing.T) {
		url := conf.Crawl["estate"].Url
		cssPath := conf.Crawl["estate"].CssPath

		rtn, err := s.crawl(url, cssPath)
		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, rtn)
		assert.Equal(t, "예정지구 지정", rtn)
		t.Log(rtn)
	})
}

func TestExchangeRate(t *testing.T) {

	s := Scraper{}
	exrate := s.ExchageRate()
	t.Log(exrate)
}

func TestFearGreedIndex(t *testing.T) {

	conf, _ := config.NewConfig()
	s := NewScraper(conf)

	rtn, err := s.FearGreedIndex()
	if err != nil {
		t.Error(err)
	}
	t.Logf("\n%+v", rtn)
}
