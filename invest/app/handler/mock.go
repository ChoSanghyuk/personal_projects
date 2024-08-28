package handler

import (
	"fmt"
	m "invest/model"
)

type AssetRetrieverMock struct {
}

func (m AssetRetrieverMock) RetrieveAssetList() ([]map[string]interface{}, error) {

	fmt.Println("RetrieveAssetList Called")
	return nil, nil
}
func (m AssetRetrieverMock) RetrieveAsset(id uint) (*m.Asset, error) {
	fmt.Println("RetrieveAsset Called")
	return nil, nil
}
func (m AssetRetrieverMock) RetrieveAssetHist(id uint) ([]m.Invest, error) {
	fmt.Println("RetrieveAssetHist Called")
	return nil, nil
}

//go:generate mockery --name AssetInfoSaver --case underscore --inpackage
type AssetInfoSaverMock struct {
}

func (m AssetInfoSaverMock) SaveAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error {
	fmt.Println("SaveAssetInfo Called")
	return nil
}
func (m AssetInfoSaverMock) UpdateAssetInfo(name string, category uint, currency string, top float64, bottom float64, selPrice float64, buyPrice float64, path string) error {
	fmt.Println("UpdateAssetInfo Called")
	return nil
}
func (m AssetInfoSaverMock) DeleteAssetInfo(id uint) error {
	fmt.Println("DeleteAssetInfo Called")
	return nil
}
