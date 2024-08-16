package handler

type FundRetriever interface {
	RetrieveFundAmount() (any, error)
	RetrieveFundAmountById(id uint) (any, error)
	RetreiveAssetOfFund() (any, error)
	RetreiveAssetOfFundById(id uint) (any, error)
}

type AssetRetriever interface {
	RetrieveAssetList() (any, error)
	RetrieveAssetInfo(id uint) (any, error)
	RetrieveAssetAmount(id uint) (any, error)
	RetrieveAssetHist(id uint) (any, error)
}

type MaketRetriever interface {
	RetrieveMarketSituation(date string) (any, error)
}

type InvestRetriever interface {
	RetrieveInvestHist(fundId string, assetId string, start string, end string) (any, error)
}

type InvestSaver interface {
	SaveHist(map[string]any) error
}

type AssetInfoSaver interface {
	SaveAssetInfo(name string, division string, peak float64, bottom float64) error
}
