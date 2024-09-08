package event

import (
	m "invest/model"
)

type Storage interface {
	RetrieveMarketStatus(date string) (*m.Market, error)
	RetrieveAssetList() ([]m.Asset, error)
	RetrieveAsset(id uint) (*m.Asset, error)
	RetreiveFundsSummaryOrderByFundId() ([]m.InvestSummary, error)
	UpdateInvestSummarySum(fundId uint, assetId uint, sum float64) error
}

type Scraper interface {
	CurrentPrice(category m.Category, code string) (float64, error)
	RealEstateStatus() (string, error)
	ExchageRate() float64
}
