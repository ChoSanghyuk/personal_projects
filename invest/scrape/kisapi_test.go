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
		WithToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0b2tlbiIsImF1ZCI6IjAyYjI2MThkLTVkMjgtNDE1Ni1hMGVkLTZjODQxZGRiYmVhMyIsInByZHRfY2QiOiIiLCJpc3MiOiJ1bm9ndyIsImV4cCI6MTcyODA4NDEzMCwiaWF0IjoxNzI3OTk3NzMwLCJqdGkiOiJQU2c0QjA4aFFDdzlqTUVvVmRYdGw1b2o1REhJaUpkY2lCZnMifQ.EFeGpl2PJ8md5f7Tgvj9Vw4QysaNpv1EmdwPeqcXD0gcroTc5nVMnAuBUqfeL3kld4HVuVLUEZoj66xiuw7nJg"),
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
		t.Log(stock.pp, stock.op, stock.hp, stock.lp, stock.ap)
	})

	t.Run("Foreign stock", func(t *testing.T) {
		pp, cp, err := s.kisForeignPrice("NAS-MSFT")
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

	t.Run("Domestic ETF", func(t *testing.T) {
		stock, err := s.kisDomesticEtfPrice("360750")
		if err != nil {
			t.Error(err)
		}
		t.Log(stock.pp, stock.op, stock.hp, stock.lp, stock.ap)
	})

	t.Run("Foreign ETF", func(t *testing.T) {
		pp, cp, err := s.kisForeignPrice("AMS-SPY")
		if err != nil {
			t.Error(err)
		}
		t.Log(pp, cp)
	})

}
