package handler

import (
	"invest/model"
	m "invest/model"
)

type FundRetriever interface {
	RetreiveFundsSummaryOrderByFundId() ([]m.InvestSummary, error)
	RetreiveFundSummaryByFundId(id uint) ([]m.InvestSummary, error)
	RetreiveAFundInvestsById(id uint) ([]m.Invest, error)
}

type FundWriter interface {
	SaveFund(name string) error
}

type AssetRetriever interface {
	RetrieveAssetList() ([]m.Asset, error)
	RetrieveAsset(id uint) (*m.Asset, error)
	RetrieveAssetHist(id uint) ([]m.Invest, error)
}

type AssetInfoSaver interface {
	SaveAssetInfo(name string, category model.Category, code string, currency string, top float64, bottom float64, selPrice float64, buyPrice float64) error
	UpdateAssetInfo(name string, category model.Category, code string, currency string, top float64, bottom float64, selPrice float64, buyPrice float64) error
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
	UpdateInvestSummaryCount(fundId uint, assetId uint, change int) error
}

type ExchageRateGetter interface {
	ExchageRate() float64
}
