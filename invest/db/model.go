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
	Id     uint
	Name   string
	Amount uint
}

type Asset struct {
	Id           uint
	Name         string
	Division     string
	Volatility   uint
	Peak         float64
	RecentPeak   float64
	RecentBottom float64
}

type FundStatus struct {
	Id      uint
	FundId  uint
	AssetId uint
	Amount  uint
}

type InvestHistory struct {
	Id       uint
	FundId   uint
	AssetId  uint
	Price    float64
	Currency string
	Count    uint
}

type Market struct {
	CreatedAt    datatypes.Date `gorm:"primaryKey"`
	Nasdaq       float64
	GreedFearIdx uint
	IndexID      uint
	PreCliIdx    CliIdx `gorm:"foreignKey:IndexID"`
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
