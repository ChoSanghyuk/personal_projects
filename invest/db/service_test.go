package db

import (
	"log"
	"testing"
)

var stg *Storage

func init() {

	dsn := "root:root@tcp(127.0.0.1:3300)/investdb?charset=utf8mb4&parseTime=True&loc=Local"
	s, err := NewStorage(dsn)
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

	// tx := stg.db.Begin()
	// rtn, err := stg.RetreiveFundsSummary()
	// if err != nil {
	// 	t.Error(t)
	// }
	// t.Log(rtn)
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

}
func TestUpdateAssetInfo(t *testing.T) {

}
func TestDeleteAssetInfo(t *testing.T) {

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

}
