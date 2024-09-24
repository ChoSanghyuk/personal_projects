package event

import (
	m "invest/model"
	"strings"
	"testing"
)

func TestEventbuySellMsg(t *testing.T) {

	stg := &StorageMock{}
	scrp := &RtPollerMock{}
	dp := &DailyPollerMock{}

	evt := NewEvent(stg, scrp, dp)

	pm := make(map[uint]float64)

	t.Run("buySellMsgTest-Buy", func(t *testing.T) {
		stg.assets = []m.Asset{
			{ID: 1, Name: "종목1", Category: m.DomesticStock, Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
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
			{ID: 1, Name: "종목1", Category: m.DomesticStock, Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
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
			{ID: 1, Name: "종목1", Category: m.DomesticStock, Code: "code", Currency: "WON", SellPrice: 480, BuyPrice: 450},
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
	scrp := &RtPollerMock{
		price: make(map[string][4]float64),
	}
	dp := &DailyPollerMock{}

	evt := NewEvent(stg, scrp, dp)

	/*
		매도 필요상황
		현재가 > ap, hp 인 애들이 앞으로 오는지
		1번 cp > ap, cp >hp
		2번 cp > ap, cp < hp (갭은 같게)
		3번 cp = ap, cp = hp
		4번 cp < ap, cp < hp
		5번 안전 자산들
	*/
	t.Run("portfolioMsg-alertwithpriority", func(t *testing.T) {

		stg.market = &m.Market{
			Status: 3,
		}
		ivsmLi := []m.InvestSummary{
			{ID: 1, FundID: 1, AssetID: 1, Asset: m.Asset{ID: 1, Category: m.Gold, Name: "금", Currency: "WON", Code: "1"}, Count: 10, Sum: 50000},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 2, Category: m.DomesticStock, Name: "삼성전자", Currency: "WON", Code: "2"}, Count: 10, Sum: 100000},
			{ID: 3, FundID: 1, AssetID: 3, Asset: m.Asset{ID: 3, Category: m.ForeignStock, Name: "애플", Currency: "USD", Code: "3"}, Count: 15, Sum: 1500},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 4, Category: m.DomesticStock, Name: "하이닉스", Currency: "WON", Code: "4"}, Count: 10, Sum: 100000},
			{ID: 2, FundID: 1, AssetID: 2, Asset: m.Asset{ID: 5, Category: m.DomesticCoin, Name: "비트코인", Currency: "WON", Code: "5"}, Count: 10, Sum: 100000},
		}
		scrp.price["5"] = [4]float64{1000, 900, 900, 0}
		scrp.price["3"] = [4]float64{1000, 900, 1100, 0}
		scrp.price["4"] = [4]float64{1000, 1000, 1000, 0}
		scrp.price["2"] = [4]float64{1000, 1100, 1100, 0}
		scrp.price["1"] = [4]float64{1000, 900, 900, 0}

		msg, err := evt.portfolioMsg(ivsmLi)

		if err != nil {
			t.Error(err)
		}
		if msg != "" {
			t.Log("\n", msg)
		} else {
			t.Error(msg)
		}
	})

}
