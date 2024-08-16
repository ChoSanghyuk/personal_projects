package handler

//go:generate mockery --name FundRetriever --case underscore --inpackage
type FundRetriever interface {
	RetrieveFundAmount() (any, error)
	RetrieveFundAmountById(id uint) (any, error)
	RetreiveAssetOfFund() (any, error)
	RetreiveAssetOfFundById(id uint) (any, error)
}

//go:generate mockery --name AssetRetriever --case underscore --inpackage
type AssetRetriever interface {
	RetrieveAssetList() (any, error)
	RetrieveAssetInfo(id uint) (any, error)
	RetrieveAssetAmount(id uint) (any, error)
	RetrieveAssetHist(id uint) (any, error)
}

//go:generate mockery --name MaketRetriever --case underscore --inpackage
type MaketRetriever interface {
	RetrieveMarketSituation(date string) (any, error)
}

//go:generate mockery --name InvestRetriever --case underscore --inpackage
type InvestRetriever interface {
	RetrieveInvestHist(fundId uint, assetId uint, start string, end string) (any, error)
}

//go:generate mockery --name InvestSaver --case underscore --inpackage
type InvestSaver interface {
	SaveInvest(fundId uint, assetId uint, price float64, currency string, count uint) error
}

//go:generate mockery --name AssetInfoSaver --case underscore --inpackage
type AssetInfoSaver interface {
	SaveAssetInfo(name string, division string, peak float64, bottom float64) error
}
