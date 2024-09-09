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
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6ImU0YmNhYzFkLTQyYTUtNDE0MC1hMjQ4LWI4NDZjYzEzMjljYiIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyNjAwOTU1OSwiaWF0IjoxNzI1OTIzMTU5LCJqdGkiOiJQU1htMG5xSzRHbUxpUlVqWWIxRFVUWG5neWxkT1JsWVdFRDAifQ.z0KUcT5qUJd2JFCHT8n1iqkw0NOI0kXBu9lgEgSOpAAr34cRobMNUfmh0hn8MI_ZTQy5j5oTVnrC6br3joWnsQ"),
	)

	t.Run("Token Generate", func(t *testing.T) {
		token, err := s.KisToken()
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})

	t.Run("Stock current Price", func(t *testing.T) {
		cp, hp, lp, err := s.kisDomesticStockPrice("005930")
		if err != nil {
			t.Error(err)
		}
		t.Log(cp, hp, lp)
	})

}
