package handler

import m "invest/model"

//go:generate mockery --name FundRetriever --case underscore --inpackage
type FundRetriever interface {
	RetrieveFundAmount() ([]m.Fund, error)
	RetreiveInvestHistOfFundById(id uint) (*m.Fund, error)
	RetreiveInvestHistOfFunds() ([]m.Fund, error)
	RetreiveAssetOfFundById(id uint) (*m.Fund, error)
}

//go:generate mockery --name AssetRetriever --case underscore --inpackage
type AssetRetriever interface {
	RetrieveAssetList() ([]map[string]interface{}, error)
	RetrieveAssetInfo(id uint) (*m.Asset, error)
	// RetrieveAssetAmount(id uint) (any, error)
	RetrieveAssetHist(id uint) ([]m.InvestHistory, error)
}

//go:generate mockery --name MaketRetriever --case underscore --inpackage
type MaketRetriever interface {
	RetrieveMarketSituation(date string) (*m.Market, error)
}

//go:generate mockery --name InvestRetriever --case underscore --inpackage
type InvestRetriever interface {
	RetrieveInvestHist(fundId uint, assetId uint, start string, end string) ([]m.InvestHistory, error)
}

//go:generate mockery --name InvestSaver --case underscore --inpackage
type InvestSaver interface {
	SaveInvest(fundId uint, assetId uint, price float64, count uint) error
}

//go:generate mockery --name AssetInfoSaver --case underscore --inpackage
type AssetInfoSaver interface {
	SaveAssetInfo(name string, division string, volatility uint, currency string, peak float64, recentPeak float64, bottom float64) error
}
