package db

import (
	m "invest/model"
	"log"
	"testing"

	"gorm.io/gorm"
)

var stg *Storage

func init() {

	dsn := "root:root@tcp(127.0.0.1:3300)/investdb?charset=utf8mb4&parseTime=True&loc=Local"
	s, err := NewStorage(dsn, &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	stg = s
}

func TestRetreiveFundsSummary(t *testing.T) {
	rtn, err := stg.RetreiveFundsSummaryOrderByFundId()
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}

func TestRetreiveFundSummaryById(t *testing.T) {

	rtn, err := stg.RetreiveFundSummaryByFundId(1)
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestRetreiveAFundInvestsById(t *testing.T) {

	rtn, err := stg.RetreiveAFundInvestsById(1)
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestRetreiveInvestHistOfFundById(t *testing.T) {

	rtn, err := stg.RetreiveInvestHistOfFundById(1)
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestSaveFund(t *testing.T) {

	err := stg.SaveFund("테스트")
	if err != nil {
		t.Error(err)
	}

	var fund m.Fund

	stg.db.Last(&fund)
	t.Logf("%+v", fund)

	if fund.Name != "테스트" {
		t.Error()
	}
	stg.db.Rollback()
}

func TestRetrieveAssetList(t *testing.T) {

	rtn, err := stg.RetrieveAssetList()
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestRetrieveAsset(t *testing.T) {

	rtn, err := stg.RetrieveAsset(1)
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestRetrieveAssetHist(t *testing.T) {

	rtn, err := stg.RetrieveAssetHist(1)
	if err != nil {
		t.Error(t)
	}
	t.Log(rtn)
}
func TestSaveAssetInfo(t *testing.T) {

	err := stg.SaveAssetInfo("테스트", m.DomesticStock, "test", "WON", 82300, 60000, 80000, 62300)
	if err != nil {
		t.Error(err)
	}

	var asset m.Asset

	stg.db.Last(&asset)
	t.Logf("%+v", asset)

	if asset.Name != "테스트" {
		t.Error()
	}
	stg.db.Rollback()

}
func TestUpdateAssetInfo(t *testing.T) {

}
func TestDeleteAssetInfo(t *testing.T) {

	err := stg.SaveAssetInfo("테스트", m.DomesticStock, "test", "WON", 82300, 60000, 80000, 62300)
	if err != nil {
		t.Error(err)
	}

	var asset m.Asset
	stg.db.Last(&asset)
	t.Logf("%+v", asset)

	err = stg.DeleteAssetInfo(asset.ID)
	if err != nil {
		t.Error(err)
	}

	stg.db.Select(&asset, asset.ID)
	t.Logf("%+v", asset)

	if asset.Name != "" {
		t.Error()
	}
	stg.db.Rollback()

}
func TestRetrieveMarketStatus(t *testing.T) {

	t.Run("날짜 미지정", func(t *testing.T) {
		rtn, err := stg.RetrieveMarketStatus("")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn)
	})

	t.Run("날짜 지정", func(t *testing.T) {
		rtn, err := stg.RetrieveMarketStatus("2024-08-29")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn)
	})

}
func TestRetrieveMarketIndicator(t *testing.T) {

	t.Run("날짜 미지정", func(t *testing.T) {
		rtn1, rtn2, err := stg.RetrieveMarketIndicator("")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn1)
		t.Log(rtn2)
	})

	t.Run("날짜 지정", func(t *testing.T) {
		rtn1, rtn2, err := stg.RetrieveMarketIndicator("2024-08-29")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn1)
		t.Log(rtn2)
	})
}
func TestSaveMarketStatus(t *testing.T) {

	err := stg.SaveMarketStatus(3)
	if err != nil {
		t.Error(err)
	}

	var mk m.Market

	stg.db.Last(&mk)
	t.Logf("%+v", mk)

	if mk.Status != 3 {
		t.Error()
	}

	stg.db.Rollback()

}
func TestRetrieveInvestHist(t *testing.T) {

	t.Run("날짜 미지정", func(t *testing.T) {
		rtn, err := stg.RetrieveInvestHist(1, 0, "", "")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn)
	})

	t.Run("날짜 지정", func(t *testing.T) {
		rtn, err := stg.RetrieveInvestHist(1, 0, "2024-05-01", "")
		if err != nil {
			t.Error(t)
		}
		t.Log(rtn)
	})
}

func TestSaveInvest(t *testing.T) {

	err := stg.SaveInvest(1, 1, 62000, 10)
	if err != nil {
		t.Error(err)
	}

	var invest m.Invest

	stg.db.Last(&invest)
	t.Logf("%+v", invest)

	if invest.Count != 10 {
		t.Error()
	}

	stg.db.Rollback()

}
