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
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6IjZmMWYyZDI5LWFhNDgtNGFkMi04M2MwLTRkMmQ4OTUzOWNjNyIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNzMwNjUzOSwiaWF0IjoxNzI3MjIwMTM5LCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.o9EH8SVSVg04R5SN4SUznjhMmYWLZ3Tsul8jTeqhScSO_sO9PVBQpJPQe7G1NEqHNartokOnf2o0PlXm6gsKwA"),
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
		t.Log(stock.pp, stock.ap, stock.hp, stock.lp)
	})

	t.Run("Foreign Index", func(t *testing.T) {
		pp, err := s.kisNasdaqIndex()
		if err != nil {
			t.Error(err)
		}
		t.Log(pp)
	})

}
