package scrape

import (
	"invest/config"
	"testing"
)

func TestKis(t *testing.T) {

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	s := NewScraper(
		conf,
		WithKIS(conf.KisAppKey(), conf.KisAppSecret()),
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6ImM0ODQ1MGEyLTc3ZjAtNDBlMC1hNjgzLTUyZGFiNzQ5MjlkOSIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNzM5MzgwNywiaWF0IjoxNzI3MzA3NDA3LCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.rU_9EzZqUp0pTywlActJLJvU5PeXTFOLhSRY7mELM6KAmDhXnuNb32nUhjFBPXxwsmfZbTBg3B32z2HZ7ZFB7Q"),
	)

	t.Run("Token Generate", func(t *testing.T) {
		token, err := s.KisToken()
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})

	t.Run("Stock current Price", func(t *testing.T) {
		stock, err := s.kisDomesticStockPrice("M04020000")
		if err != nil {
			t.Error(err)
		}
		t.Log(stock.pp, stock.hp, stock.lp)
	})

	t.Run("Foreign stock", func(t *testing.T) {
		pp, cp, err := s.kisForeignStockPrice("NAS-MSFT")
		if err != nil {
			t.Error(err)
		}
		t.Log(pp, cp, err)
	})

	t.Run("Foreign Index", func(t *testing.T) {
		pp, err := s.kisNasdaqIndex()
		if err != nil {
			t.Error(err)
		}
		t.Log(pp)
	})

}
