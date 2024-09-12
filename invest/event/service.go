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

type RtPoller interface {
	CurrentPrice(category m.Category, code string) (float64, error)
	RealEstateStatus() (string, error)
}

type DailyPoller interface {
	ExchageRate() float64
	FearGreedIndex() (uint, error)
	Nasdaq() (float64, error)
	CliIdx() (float64, error)
}
