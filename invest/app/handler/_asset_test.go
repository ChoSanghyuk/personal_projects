package handler

import (
	"invest/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestAssetGetHandler(t *testing.T) {

	app := fiber.New()

	readerMock := NewMockAssetRetriever(t)
	setAssetRetrieverMock(readerMock)

	f := AssetHandler{
		r: readerMock,
		w: nil,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("Assets", func(t *testing.T) {
		err := sendReqeust(app, "/assets", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("AssetInfo", func(t *testing.T) {
		err := sendReqeust(app, "/assets/1", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("AssetAmount", func(t *testing.T) {
		err := sendReqeust(app, "/assets/1/amount", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("AssetHist", func(t *testing.T) {
		err := sendReqeust(app, "/assets/1/hist", "GET", nil)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

func TestAssetPostHandler(t *testing.T) {

	app := fiber.New()

	writerMock := NewMockAssetInfoSaver(t)
	setAssetSaverMock(writerMock)

	f := AssetHandler{
		r: nil,
		w: writerMock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("SaveAssets", func(t *testing.T) {
		reqBody := model.Asset{
			Name:         "test",
			Category:     "stock",
			Top:          500,
			RecentBottom: 400,
		}
		err := sendReqeust(app, "/assets", "POST", reqBody)
		assert.NoError(t, err)
	})

	t.Run("SaveAssets_InvalidReq", func(t *testing.T) {
		reqBody := model.Asset{
			// Name:     "test",
			Category:     "stock",
			Top:          500,
			RecentBottom: 400,
		}
		err := sendReqeust(app, "/assets", "POST", reqBody)
		assert.Error(t, err)
	})
}

func setAssetRetrieverMock(m *MockAssetRetriever) error {

	m.On("RetrieveAssetList").Return([]map[string]interface{}{{"ID": 1, "Name": "비트코인"}, {"ID": 2, "Name": "solx"}}, nil)
	m.On("RetrieveAssetInfo", mock.AnythingOfType("uint")).Return(&model.Asset{ID: 1, Name: "비트코인"}, nil)
	// m.On("RetrieveAssetAmount", mock.AnythingOfType("uint")).Return("RetrieveAssetAmount Called", nil)
	m.On("RetrieveAssetHist", mock.AnythingOfType("uint")).Return([]model.Invest{{ID: 1, FundID: 3, AssetID: 1}}, nil)

	return nil
}

func setAssetSaverMock(m *MockAssetInfoSaver) error {
	m.On("SaveAssetInfo", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("uint"), mock.AnythingOfType("string"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64"), mock.AnythingOfType("float64")).Return(nil)
	return nil
}
