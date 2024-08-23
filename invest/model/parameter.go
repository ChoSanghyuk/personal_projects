package model

type GetMarketParam struct {
	Date string `json:"date" validate:"date"`
}

type GetInvestHistParam struct {
	FundId    uint   `json:"fund_id"`
	AssetId   uint   `json:"asset_id"`
	StartDate string `json:"start_date" validate:"date"`
	EndDate   string `json:"end_date" validate:"date"`
}

type SaveInvestParam struct {
	FundId  uint    `json:"fund_id"`
	AssetId uint    `json:"asset_id"`
	Price   float64 `json:"price"`
	Count   uint    `json:"count"`
}
