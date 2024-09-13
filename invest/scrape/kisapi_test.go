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
		WithToken(""),
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

	t.Run("Foreing Index", func(t *testing.T) {
		s.kisForeingIndex()
	})

}

// 2852601
//   17569.68
//      72.57
