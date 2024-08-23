package db

import (
	m "invest/model"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func (s Storage) RetrieveFundAmount() ([]m.Fund, error) {

	var funds []m.Fund

	result := s.db.Find(&funds)
	if result.Error != nil {
		return nil, result.Error
	}

	return funds, nil
}

func (s Storage) RetrieveFundHistById(id uint) (*m.Fund, error) {

	var fund m.Fund

	result := s.db.First(&fund, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fund, nil
}

func (s Storage) RetreiveInvestHistOfFund() ([]m.Fund, error) {

	var funds []m.Fund

	result := s.db.Model(&m.Fund{}).Preload("Hist.Asset").Find(&funds)

	if result.Error != nil {
		return nil, result.Error
	}

	return funds, nil
}

func (s Storage) RetreiveAssetOfFundById(id uint) (*m.Fund, error) {

	var fund m.Fund

	result := s.db.Model(&m.Fund{}).Preload("Hist.Asset").Find(&fund, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &fund, nil
}

func (s Storage) RetrieveAssetList() ([]map[string]interface{}, error) {

	var assets []map[string]interface{}

	result := s.db.Model(&m.Asset{}).Select("id", "name").Find(&assets)
	if result.Error != nil {
		return nil, result.Error
	}
	return assets, nil
}

func (s Storage) RetrieveAssetInfo(id uint) (*m.Asset, error) {

	var asset m.Asset

	result := s.db.First(&asset, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &asset, nil
}

func (s Storage) RetrieveAssetHist(id uint) (any, error) {

	var asset m.Asset

	result := s.db.Model(&m.Asset{}).Preload("Hist").Find(&asset, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return asset.Hist, nil
}

func (s Storage) RetrieveMarketSituation(date string) (*m.Market, error) {

	var market m.Market

	result := s.db.Find(&market, date)
	if result.Error != nil {
		return nil, result.Error
	}
	return &market, nil
}

func (s Storage) RetrieveInvestHist(fundId uint, assetId uint, start string, end string) ([]m.InvestHistory, error) {

	conditionmap := map[string]interface{}{}
	if fundId != 0 {
		conditionmap["fund_id"] = fundId
	}
	if assetId != 0 {
		conditionmap["asset_id"] = assetId
	}
	if start != "" {
		conditionmap["start"] = start
	}
	if end != "" {
		conditionmap["end"] = end
	}

	var investHist []m.InvestHistory

	result := s.db.Where(conditionmap).Find(&investHist)
	if result.Error != nil {
		return nil, result.Error
	}

	return investHist, nil
}

func (s Storage) SaveInvest(fundId uint, assetId uint, price float64, currency string, count uint) error {

	result := s.db.Create(&m.InvestHistory{
		FundID:       fundId,
		AssetID:      assetId,
		CurrentPrice: price,
		Count:        count,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) SaveAssetInfo(name string, division string, volatility uint, currency string, peak float64, recentPeak float64, bottom float64) error {

	result := s.db.Create(&m.Asset{
		Name:         name,
		Division:     division,
		Volatility:   volatility,
		Currency:     currency,
		Peak:         peak,
		RecentPeak:   recentPeak,
		RecentBottom: bottom,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
