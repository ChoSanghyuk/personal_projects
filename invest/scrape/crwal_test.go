package scrape

import (
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoldCrwal(t *testing.T) {

	s := Scraper{}

	url := config.ConfigInfo.Gold.Crawl.Url
	cssPath := config.ConfigInfo.Gold.Crawl.CssPath
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

	t.Run("Crwal", func(t *testing.T) {
		url := config.ConfigInfo.Bitcoin.Crawl.Url
		cssPath := config.ConfigInfo.Bitcoin.Crawl.CssPath

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

	t.Run("Crwal", func(t *testing.T) {
		url := config.ConfigInfo.RealEstate.Crawl.Url
		cssPath := config.ConfigInfo.RealEstate.Crawl.CssPath

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
