package handler

import m "invest/model"

//go:generate mockery --name FundRetriever --case underscore --inpackage
type FundRetriever interface {
	RetreiveFundsSummary() ([]m.InvestSummary, error)
	RetreiveFundSummaryById(id uint) ([]m.InvestSummary, error)
	RetreiveAFundInvestsById(id uint) ([]m.Invest, error)

	RetrieveFundAmount() ([]m.Fund, error)
	RetreiveInvestHistOfFundById(id uint) (*m.Fund, error)
}

type FundWriter interface {
	SaveFund(name string) error
}

//go:generate mockery --name AssetRetriever --case underscore --inpackage
type AssetRetriever interface {
	RetrieveAssetList() ([]map[string]interface{}, error)
	RetrieveAssetInfo(id uint) (*m.Asset, error)
	// RetrieveAssetAmount(id uint) (any, error)
	RetrieveAssetHist(id uint) ([]m.Invest, error)
}

//go:generate mockery --name MaketRetriever --case underscore --inpackage
type MaketRetriever interface {
	RetrieveMarketSituation(date string) (*m.Market, error)
}

//go:generate mockery --name InvestRetriever --case underscore --inpackage
type InvestRetriever interface {
	RetrieveInvestHist(fundId uint, assetId uint, start string, end string) ([]m.Invest, error)
}

//go:generate mockery --name InvestSaver --case underscore --inpackage
type InvestSaver interface {
	SaveInvest(fundId uint, assetId uint, price float64, count uint) error
}

//go:generate mockery --name AssetInfoSaver --case underscore --inpackage
type AssetInfoSaver interface {
	SaveAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error
	UpdateAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error
	DeleteAssetInfo(id uint) error
}
