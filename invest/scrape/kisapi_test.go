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
		WithKIS(conf.KisAppKey(), conf.KisAppSecret()),
		// WithToken(""),
	)

	t.Run("Token Generate", func(t *testing.T) {
		token, err := s.KisToken()
		if err != nil {
			t.Error(err)
		}
		t.Log(token)
	})

	t.Run("Stock current Price", func(t *testing.T) {
		price, err := s.KisCurrentPrice("")
		if err != nil {
			t.Error(err)
		}
		t.Log(price)
	})

}
