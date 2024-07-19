package crawl

import (
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGold(t *testing.T) {

	url := config.ConfigInfo.Gold.Crawl.Url
	cssPath := config.ConfigInfo.Gold.Crawl.CssPath
	// CallGoldApi()
	rtn, err := Crawl(url, cssPath)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}

func TestBitcoin(t *testing.T) {

	t.Run("Crwal", func(t *testing.T) {
		url := config.ConfigInfo.Bitcoin.Crawl.Url
		cssPath := config.ConfigInfo.Bitcoin.Crawl.CssPath
		// CallGoldApi()
		rtn, err := Crawl(url, cssPath)
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
