package handler

import (
	model "invest/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

func TestFundHandler(t *testing.T) {

	app := fiber.New()

	mock := NewMockFundRetriever(t)
	setFundRetrieverMock(mock)

	f := FundHandler{
		r: mock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("TotalFunds", func(t *testing.T) {
		err := sendReqeust(app, "/funds", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("TotalFundAssets", func(t *testing.T) {
		err := sendReqeust(app, "/funds/assets", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("FundAsset", func(t *testing.T) {
		err := sendReqeust(app, "/funds/1/assets", "GET", nil)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

/*
************************************************
Mock
*************************************************
*/

func setFundRetrieverMock(m *MockFundRetriever) error {

	m.On("RetreiveAssetOfFundById", mock.AnythingOfType("uint")).Return(&model.Fund{ID: 3, Name: "개인"}, nil)
	m.On("RetreiveInvestHistOfFund").Return([]model.Fund{}, nil)
	m.On("RetrieveFundAmount").Return([]model.Fund{{ID: 3, Name: "개인"}}, nil)
	m.On("RetreiveInvestHistOfFundById", mock.AnythingOfType("uint")).Return(&model.Fund{}, nil)

	return nil
}
