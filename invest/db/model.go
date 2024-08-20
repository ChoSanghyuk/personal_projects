package db

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
	Id          uint
	Name        string
	MarketValue uint
}

type Asset struct {
	Id           uint
	Name         string
	Division     string
	Volatility   uint
	Currency     string
	Peak         float64
	RecentPeak   float64
	RecentBottom float64
}

type InvestHistory struct {
	Id           uint
	FundId       Fund  `gorm:"foreignKey:FundID"`
	AssetId      Asset `gorm:"foreignKey:AssetID"`
	CurrentPrice float64
	Count        uint
	gorm.Model
}

type Market struct {
	CreatedAt    datatypes.Date `gorm:"primaryKey"`
	Nasdaq       float64
	GreedFearIdx uint
}

type CliIdx struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt datatypes.Date
	Idx       float64
}

type Sample struct {
	Id   uint `gorm:"primaryKey"`
	Date datatypes.Date
	Time time.Time
	gorm.Model
}
