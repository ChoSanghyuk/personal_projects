package model

type GetInvestHistParam struct {
	FundId    uint   `json:"fund_id"`
	AssetId   uint   `json:"asset_id"`
	StartDate string `json:"start_date" validate:"date"`
	EndDate   string `json:"end_date" validate:"date"`
}
