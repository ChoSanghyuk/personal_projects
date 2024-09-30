package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestAssetHandler(t *testing.T) {

	app := fiber.New()

	readerMock := NewMockAssetRetriever(t)
	setAssetRetrieverMock(readerMock)
	writerMock := NewMockAssetInfoSaver(t)
	setAssetSaverMock(writerMock)

	f := AssetHandler{
		r: readerMock,
		w: writerMock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("TotalFunds", func(t *testing.T) {
		err := sendReqeust(app, "/funds", nil)
		assert.NoError(t, err)
	})

	t.Run("Fund", func(t *testing.T) {
		err := sendReqeust(app, "/funds/1", nil)
		assert.NoError(t, err)
	})

	t.Run("TotalFundAssets", func(t *testing.T) {
		err := sendReqeust(app, "/funds/assets", nil)
		assert.NoError(t, err)
	})

	t.Run("FundAsset", func(t *testing.T) {
		err := sendReqeust(app, "/funds/1/assets", nil)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

func setAssetRetrieverMock(m *MockAssetRetriever) error {

	m.On("RetrieveAssetAmount").Return("hello1", nil)
	m.On("RetrieveAssetHist", mock.AnythingOfType("uint")).Return("hello2", nil)
	m.On("RetrieveAssetInfo").Return("hello3", nil)
	m.On("RetrieveAssetList", mock.AnythingOfType("uint")).Return("hello4", nil)

	return nil
}

func setAssetSaverMock(m *MockAssetInfoSaver) error {
	m.On("SaveAssetInfo").Return("hello1", nil)
	return nil
}
