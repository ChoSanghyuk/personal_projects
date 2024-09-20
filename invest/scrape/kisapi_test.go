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
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6Ijk3ZDRiNjRhLThlYTktNDBiOC04MjRjLTEwMjJjYjVjZTc0OSIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNjg3NDEzMywiaWF0IjoxNzI2Nzg3NzMzLCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.2PoEX5DZ1LyTTq4JX0S-faY3S7i5qxESIStdP76rWW3XB7NK_0WGXDbYPrF6Ss3WGydYnaGX15r1_UYucaj3VA"),
	)

	t.Run("Token Generate", func(t *testing.T) {
		token, err := s.KisToken()
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})

	t.Run("Stock current Price", func(t *testing.T) {
		cp, hp, lp, err := s.kisDomesticStockPrice("M04020000")
		if err != nil {
			t.Error(err)
		}
		t.Log(cp, hp, lp)
	})

	t.Run("Foreign Index", func(t *testing.T) {
		cp, err := s.kisNasdaqIndex()
		if err != nil {
			t.Error(err)
		}
		t.Log(cp)
	})

}
