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
	Name         string
	Division     string
	Volatility   uint
	Currency     string
	Peak         float64
	RecentPeak   float64
	RecentBottom float64
	Hist         []InvestHistory
}

type InvestHistory struct {
	ID           uint
	FundID       uint
	AssetID      uint
	Asset        Asset
	CurrentPrice float64
	Count        uint
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
