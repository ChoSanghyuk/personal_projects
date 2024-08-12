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
	RetrieveMarketSituation() (any, error)
}

type InvestSaver interface {
	SaveHist(map[string]any) (any, error)
}

type AssetInfoSaver interface {
	SaveAssetInfo(map[string]any) (any, error)
}

type MarketInfoSaver interface {
	SaveMakreSitutation() (any, error)
}
