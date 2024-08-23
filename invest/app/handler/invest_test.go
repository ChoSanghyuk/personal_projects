package handler

import (
	"invest/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInvestGetHandler(t *testing.T) {

	app := fiber.New()

	m := NewMockInvestRetriever(t)
	setInvestRetrieverMock(m)

	f := InvestHandler{
		r: m,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("InvestHist", func(t *testing.T) {
		reqBody := model.GetInvestHistParam{
			FundId:    1,
			AssetId:   1,
			StartDate: "2024-08-10",
			EndDate:   "2024-08-20",
		}
		err := sendReqeust(app, "/invest/hist", "GET", reqBody)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

func TestInvestPostHandler(t *testing.T) {

	app := fiber.New()

	m := NewMockInvestSaver(t)
	setInvestSaverMock(m)

	f := InvestHandler{
		w: m,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("SaveInvest", func(t *testing.T) {
		reqBody := model.SaveInvestParam{
			FundId:  1,
			AssetId: 1,
			Price:   100,
			Count:   5,
		}
		err := sendReqeust(app, "/invest", "POST", reqBody)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

func setInvestRetrieverMock(m *MockInvestRetriever) {
	m.On("RetrieveInvestHist", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]model.InvestHistory{}, nil)
}

func setInvestSaverMock(m *MockInvestSaver) {
	m.On("SaveInvest", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("float64"), mock.AnythingOfType("uint")).Return(nil)
}
