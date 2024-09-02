package scrape

import (
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoldCrwal(t *testing.T) {

	s := Scraper{}
	conf, _ := config.NewConfig()

	url := conf.Crawl["gold"].Url
	cssPath := conf.Crawl["gold"].CssPath
	// CallGoldApi()
	rtn, err := s.Crawl(url, cssPath)
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

		// CallGoldApi()
		rtn, err := s.Crawl(url, cssPath)
		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, rtn)

		t.Log(rtn)
	})

	// t.Run("Crwal From Chrome", func(t *testing.T) {

	// 	url := config.ConfigInfo.Bitcoin.Crawl.Url
	// 	cssPath := config.ConfigInfo.Bitcoin.Crawl.CssPath

	// 	// CallGoldApi()
	// 	rtn, err := CrawlChrome(url, cssPath)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}

	// 	assert.NotEmpty(t, rtn)

	// 	t.Log(rtn)
	// })

}

func TestEstateCrwal(t *testing.T) {

	s := Scraper{}
	conf, _ := config.NewConfig()

	t.Run("Crwal", func(t *testing.T) {
		url := conf.Crawl["estate"].Url
		cssPath := conf.Crawl["estate"].CssPath

		// CallGoldApi()
		rtn, err := s.Crawl(url, cssPath)
		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, rtn)
		assert.Equal(t, "예정지구 지정", rtn)
		t.Log(rtn)
	})
}
