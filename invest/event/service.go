package event

import m "invest/model"

type Storage interface {
	RetrieveMarketStatus(date string) (*m.Market, error)
	RetrieveAssetList() ([]m.Asset, error)
	RetrieveAsset(id uint) (*m.Asset, error)
	RetreiveFundsSummaryOrderByFundId() ([]m.InvestSummary, error)
	UpdateInvestSummarySum(fundId uint, assetId uint, sum float64) error
}

type Scraper interface {
	CallApi(url string, header map[string]string) (string, error)
	Crawl(url string, cssPath string) (string, error)
	GetRealtimeExchageRate() float64
}

type Transmitter interface {
	ApiInfo(target string) (url string, header map[string]string)
	CrawlInfo(target string) (url string, cssPath string)
}
