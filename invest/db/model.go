package db

// type DivisionCode uint

// const (
// 	cash DivisionCode = iota+1
// 	growth
// 	dividend

// )

type DangerLevel uint

const (
	no DangerLevel = iota + 1
	low
	moderate
	high
	extremeHigh
)

type Fund struct {
	Id     uint
	Name   string
	Amount uint
}

type Asset struct {
	Id       uint
	Name     string
	Division string
	// DangerLevel DangerLevel
	Peak   float64
	Bottom float64
}

type FundStatus struct {
	Id      uint
	FundId  uint
	AssetId uint
	Amount  uint
}

type InvestHistory struct {
	Id      uint
	FundId  uint
	AssetId uint
	Price   float64
	Count   uint
}
