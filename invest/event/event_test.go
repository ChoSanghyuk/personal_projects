package event

import (
	m "invest/model"
	"strings"
	"testing"
)

func TestEventbuySellMsg(t *testing.T) {

	stg := &StorageMock{}
	scrp := &ScraperMock{}

	evt := NewEvent(stg, scrp)

	pm := make(map[uint]float64)

	t.Run("buySellMsgTest-Buy", func(t *testing.T) {
		stg.assets = []m.Asset{
			{ID: 1, Name: "종목1", Category: "국내주식", Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
		}
		scrp.cp = 400
		msg, err := evt.buySellMsg(1, pm)
		if err != nil {
			t.Error(err)
		}
		if strings.Contains(msg, "BUY") {
			t.Log(msg)
		} else {
			t.Error(msg)
		}
	})

	t.Run("buySellMsgTest-Sell", func(t *testing.T) {
		stg.assets = []m.Asset{
			{ID: 1, Name: "종목1", Category: "국내주식", Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
		}
		scrp.cp = 490
		msg, err := evt.buySellMsg(1, pm)
		if err != nil {
			t.Error(err)
		}
		if strings.Contains(msg, "SELL") {
			t.Log(msg)
		} else {
			t.Error(msg)
		}
	})

	t.Run("buySellMsgTest-Nothing", func(t *testing.T) {
		stg.assets = []m.Asset{
			{ID: 1, Name: "종목1", Category: "국내주식", Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
		}
		scrp.cp = 470
		msg, err := evt.buySellMsg(1, pm)
		if err != nil {
			t.Error(err)
		}
		if msg == "" {
			t.Log(msg)
		} else {
			t.Error(msg)
		}
	})
}

func TestEventportfolioMsg(t *testing.T) {

	stg := &StorageMock{}
	scrp := &ScraperMock{}

	evt := NewEvent(stg, scrp)

	pm := make(map[uint]float64)

	t.Run("portfolioMsg-alertwithpriority", func(t *testing.T) { // todo. 현금 현황 관리

		stg.market = &m.Market{
			Status: 3,
		}
		ivsmLi := []m.InvestSummary{
			{ID: 1, FundID: 1, AssetID: 1, Asset: m.Asset{ID: 1, Category: "금", Currency: "WON", Top: 10000}, Count: 10, Sum: 50000},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 2, Category: "국내주식", Name: "삼성전자", Currency: "WON", Top: 10000}, Count: 10, Sum: 100000},
		}

		pm[1] = 9000
		pm[2] = 11000

		msg, err := evt.portfolioMsg(ivsmLi, pm)

		if err != nil {
			t.Error(err)
		}
		if msg != "" {
			t.Log("\n", msg)
		} else {
			t.Error(msg)
		}
	})

	t.Run("portfolioMsg-alertwithoutpriority", func(t *testing.T) { // todo. 현금 현황 관리

		stg.market = &m.Market{
			Status: 3,
		}
		ivsmLi := []m.InvestSummary{
			{ID: 1, FundID: 1, AssetID: 1, Asset: m.Asset{ID: 1, Category: "금", Name: "금", Currency: "WON", Top: 10000}, Count: 10, Sum: 50000},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 2, Category: "국내주식", Name: "삼성전자", Currency: "WON", Top: 10000}, Count: 10, Sum: 100000},
		}

		pm[1] = 11000
		pm[2] = 9000

		msg, err := evt.portfolioMsg(ivsmLi, pm)

		if err != nil {
			t.Error(err)
		}
		if msg != "" {
			t.Log("\n", msg)
		} else {
			t.Error(msg)
		}
	})

	t.Run("portfolioMsg-nothing", func(t *testing.T) {
		stg.market = &m.Market{
			Status: 3,
		}
		ivsmLi := []m.InvestSummary{
			{ID: 1, FundID: 1, AssetID: 1, Asset: m.Asset{ID: 1, Category: "금", Currency: "WON", Top: 10000}, Count: 10, Sum: 150000},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 2, Category: "국내주식", Name: "삼성전자", Currency: "WON", Top: 10000}, Count: 10, Sum: 100000},
		}

		pm[1] = 9000
		pm[2] = 11000

		msg, err := evt.portfolioMsg(ivsmLi, pm)

		if err != nil {
			t.Error(err)
		}
		if msg == "" {
			t.Log("\n", msg)
		} else {
			t.Error(msg)
		}
	})

}
