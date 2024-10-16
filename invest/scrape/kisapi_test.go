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
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6IjdkMmY3Y2IxLTk4NzctNGNhZS1iMWY3LTU2ZjJlN2E5OGM4OCIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNzQ4MDcwMCwiaWF0IjoxNzI3Mzk0MzAwLCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.mNbB10oe2ZtAS8ddP0UPjxMT9HB4g5wMikzk8yZXK8_wo_aDZSdUjj2O-XyDvwTdMBKzDVn-SCLDh2rPJh95mQ"),
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
		t.Log(stock.pp, stock.op, stock.hp, stock.lp)
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
