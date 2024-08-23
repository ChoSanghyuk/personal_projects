package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// type DangerLevel uint

// const (
// 	no DangerLevel = iota + 1
// 	low
// 	moderate
// 	high
// 	extremeHigh
// )

type Fund struct {
	ID          uint
	Name        string
	MarketValue uint
	// assets      []Asset
	Hist []InvestHistory
}

type Asset struct {
	ID           uint
	Name         string          `json:"name" validate:"required"`
	Division     string          `json:"division"`
	Volatility   uint            `json:"volatility"`
	Currency     string          `json:"currency"`
	Peak         float64         `json:"peak"`
	RecentPeak   float64         `json:"recent_peak"`
	RecentBottom float64         `json:"bottom"`
	Hist         []InvestHistory `json:"hist"`
}

type InvestHistory struct {
	ID      uint
	FundID  uint `json:"fund_id"`
	AssetID uint `json:"asset_id"`
	Asset   Asset
	Price   float64 `json:"price"`
	Count   uint    `json:"count"`
	gorm.Model
}

type Market struct {
	CreatedAt    datatypes.Date `gorm:"primaryKey"`
	Nasdaq       float64
	GreedFearIDx uint
}

type CliIdx struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt datatypes.Date
	IDx       float64
}

type Sample struct {
	ID   uint `gorm:"primaryKey"`
	Date datatypes.Date
	Time time.Time
	gorm.Model
}
