package handler

type SaveAssetParam struct {
	Id       uint
	Name     string `validate:"required"`
	Division string
	Peak     float64
	Bottom   float64
}

type GetMarketParam struct {
	Date string `validate:"date"`
}
