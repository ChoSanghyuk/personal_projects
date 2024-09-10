package event

import md "invest/model"

type StorageMock struct {
	market *md.Market
	assets []md.Asset
	ivsm   []md.InvestSummary
	err    error
}

func (m StorageMock) RetrieveMarketStatus(date string) (*md.Market, error) {
	if m.err != nil {
		return nil, m.err
	}

	return m.market, nil
}

func (m StorageMock) RetrieveAssetList() ([]md.Asset, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.assets, nil
}

func (m StorageMock) RetrieveAsset(id uint) (*md.Asset, error) {
	if m.err != nil {
		return nil, m.err
	}
	for _, a := range m.assets {
		if a.ID == id {
			return &a, nil
		}
	}
	return &md.Asset{}, nil
}

func (m StorageMock) RetreiveFundsSummaryOrderByFundId() ([]md.InvestSummary, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.ivsm, nil
}

func (m StorageMock) UpdateInvestSummarySum(fundId uint, assetId uint, sum float64) error {
	if m.err != nil {
		return m.err
	}
	return nil
}

type ScraperMock struct {
	cp     float64
	estate string
	err    error
}

func (m ScraperMock) CurrentPrice(category md.Category, code string) (float64, error) {
	if m.err != nil {
		return 0, m.err
	}
	return m.cp, nil
}

func (m ScraperMock) RealEstateStatus() (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.estate, nil
}

func (m ScraperMock) ExchageRate() float64 {
	if m.err != nil {
		return 0
	}

	return 1300
}
