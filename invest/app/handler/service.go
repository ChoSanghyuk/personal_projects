package handler

import m "invest/model"

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

type AssetRetriever interface {
	RetrieveAssetList() ([]map[string]interface{}, error)
	RetrieveAsset(id uint) (*m.Asset, error)
	RetrieveAssetHist(id uint) ([]m.Invest, error)
}

type AssetInfoSaver interface {
	SaveAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error
	UpdateAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error
	DeleteAssetInfo(id uint) error
}

type MaketRetriever interface {
	RetrieveMarketStatus(date string) (*m.Market, error)
	RetrieveMarketIndicator(date string) (*m.DailyIndex, *m.CliIndex, error)
}

type MarketSaver interface {
	SaveMarketStatus(status uint) error
}

type InvestRetriever interface {
	RetrieveInvestHist(fundId uint, assetId uint, start string, end string) ([]m.Invest, error)
}

type InvestSaver interface {
	SaveInvest(fundId uint, assetId uint, price float64, count int) error
}
