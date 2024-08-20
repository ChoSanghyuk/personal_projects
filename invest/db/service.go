package db

import (
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func (s Storage) RetrieveFundAmount() (any, error) {

	var funds []Fund

	result := s.db.Find(&funds)
	if result.Error != nil {
		return nil, result.Error
	}

	return funds, nil
}

func (s Storage) RetrieveFundAmountById(id uint) (any, error) {

	var fund Fund

	result := s.db.First(&fund, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return fund, nil
}

func (s Storage) RetreiveAssetOfFund() (any, error) {

	var fundsStatus []FundStatus

	result := s.db.Find(&fundsStatus)
	if result.Error != nil {
		return nil, result.Error
	}

	return fundsStatus, nil
}

func (s Storage) RetreiveAssetOfFundById(id uint) (any, error) {

	var fundStatus FundStatus

	result := s.db.Where("fund_id = ?", id).First(&fundStatus)
	if result.Error != nil {
		return nil, result.Error
	}

	return fundStatus, nil
}

func (s Storage) RetrieveAssetList() (any, error) {

	var assets []map[string]interface{}

	result := s.db.Model(&Asset{}).Select("id", "name").Find(&assets)
	if result.Error != nil {
		return nil, result.Error
	}
	return assets, nil
}

func (s Storage) RetrieveAssetInfo(id uint) (any, error) {

	var asset Asset

	result := s.db.First(&asset, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return asset, nil
}

// func (s Storage) RetrieveAssetAmount(id uint) (any, error) {

// 	var fundStatus FundStatus

// 	result := s.db.Where("asset_id = ?", id).First(&fundStatus)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return fundStatus, nil
// }

func (s Storage) RetrieveAssetHist(id uint) (any, error) {

	var investHist []InvestHistory

	result := s.db.Where("asset_id = ?", id).First(&investHist)
	if result.Error != nil {
		return nil, result.Error
	}

	return investHist, nil
}

func (s Storage) RetrieveMarketSituation(date string) (any, error) {

	var market Market

	result := s.db.Find(&market, date)
	if result.Error != nil {
		return nil, result.Error
	}
	return market, nil
}

func (s Storage) RetrieveInvestHist(fundId uint, assetId uint, start string, end string) (any, error) {

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

	var investHist []InvestHistory

	result := s.db.Where(conditionmap).Find(&investHist)
	if result.Error != nil {
		return nil, result.Error
	}

	return investHist, nil
}

func (s Storage) SaveInvest(fundId uint, assetId uint, price float64, currency string, count uint) error {

	result := s.db.Create(&InvestHistory{
		FundId:       fundId,
		AssetId:      assetId,
		CurrentPrice: price,
		Currency:     currency,
		Count:        count,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s Storage) SaveAssetInfo(name string, division string, volatility uint, currency string, peak float64, recentPeak float64, bottom float64) error {

	result := s.db.Create(&Asset{
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
