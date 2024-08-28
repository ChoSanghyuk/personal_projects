package db

import (
	m "invest/model"

	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func (s Storage) RetreiveFundsSummary() ([]m.InvestSummary, error) {

	var fundsSummary []m.InvestSummary

	result := s.db.Model(&m.InvestSummary{}).Preload("Fund").Find(&fundsSummary)

	if result.Error != nil {
		return nil, result.Error
	}

	return fundsSummary, nil

}

func (s Storage) RetreiveFundSummaryById(id uint) ([]m.InvestSummary, error) {

	var fundsSummary []m.InvestSummary

	result := s.db.Model(&m.InvestSummary{}).Preload("Fund").Find(&fundsSummary, id) // .Order("asset_id")

	if result.Error != nil {
		return nil, result.Error
	}

	return fundsSummary, nil

}

func (s Storage) RetreiveAFundInvestsById(id uint) ([]m.Invest, error) {
	var invets []m.Invest

	result := s.db.Model(&m.Invest{}).Where(&m.Invest{FundID: id}, "fund_id").Find(&invets) // .Order("asset_id")

	if result.Error != nil {
		return nil, result.Error
	}

	return invets, nil
}

func (s Storage) RetrieveFundAmount() ([]m.Fund, error) {

	var funds []m.Fund

	result := s.db.Find(&funds)
	if result.Error != nil {
		return nil, result.Error
	}

	return funds, nil
}

func (s Storage) RetreiveInvestHistOfFundById(id uint) (*m.Fund, error) {
	var fund m.Fund

	result := s.db.First(&fund, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &fund, nil
}

func (s Storage) SaveFund(name string) error {

	result := s.db.Create(&m.Fund{
		Name: name,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Storage) RetrieveAssetList() ([]map[string]interface{}, error) {

	var assets []map[string]interface{}

	result := s.db.Model(&m.Asset{}).Select("id", "name").Find(&assets)
	if result.Error != nil {
		return nil, result.Error
	}
	return assets, nil
}

func (s Storage) RetrieveAsset(id uint) (*m.Asset, error) {

	var asset m.Asset

	result := s.db.First(&asset, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &asset, nil
}

func (s Storage) RetrieveAssetHist(id uint) ([]m.Invest, error) {

	var invests []m.Invest

	result := s.db.Model(&m.Invest{}).Find(&invests, id) // Preload("Asset")
	if result.Error != nil {
		return nil, result.Error
	}

	return invests, nil
}

func (s Storage) SaveAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error {

	result := s.db.Create(&m.Asset{
		Name:      name,
		Category:  category,
		Currency:  currency,
		Top:       top,
		Bottom:    bottom,
		SellPrice: selPrice,
		BuyPrice:  buyPrice,
		Path:      path,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Storage) UpdateAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error {

	result := s.db.Updates(m.Asset{
		Name:      name,
		Category:  category,
		Currency:  currency,
		Top:       top,
		Bottom:    bottom,
		SellPrice: selPrice,
		BuyPrice:  buyPrice,
		Path:      path,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Storage) DeleteAssetInfo(id uint) error {

	result := s.db.Delete(&m.Asset{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Storage) RetrieveMarketStatus(date string) (*m.Market, error) {

	var market m.Market

	if date == "" {
		result := s.db.Last(&market) // Preload("Asset")
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result := s.db.First(&market, date) // Preload("Asset")
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return &market, nil
}

func (s Storage) RetrieveMarketIndicator(date string) (*m.DailyIndex, *m.CliIndex, error) {

	var dailyIdx m.DailyIndex
	var cliIdx m.CliIndex

	if date == "" {
		result := s.db.Last(&dailyIdx) // Preload("Asset")
		if result.Error != nil {
			return nil, nil, result.Error
		}

		result = s.db.Last(&cliIdx) // Preload("Asset")
		if result.Error != nil {
			return nil, nil, result.Error
		}
	} else {
		result := s.db.First(&dailyIdx, date) // Preload("Asset")
		if result.Error != nil {
			return nil, nil, result.Error
		}

		result = s.db.First(&cliIdx, date) // Preload("Asset")
		if result.Error != nil {
			return nil, nil, result.Error
		}
	}

	return &dailyIdx, &cliIdx, nil
}

func (s Storage) SaveMarketStatus(status uint) error {

	result := s.db.Create(&m.Market{
		Status: status,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s Storage) RetrieveInvestHist(fundId uint, assetId uint, start string, end string) ([]m.Invest, error) {

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

	var investHist []m.Invest

	result := s.db.Where(conditionmap).Find(&investHist)
	if result.Error != nil {
		return nil, result.Error
	}

	return investHist, nil
}

func (s Storage) SaveInvest(fundId uint, assetId uint, price float64, count int) error {

	result := s.db.Create(&m.Invest{
		FundID:  fundId,
		AssetID: assetId,
		Price:   price,
		Count:   count,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
