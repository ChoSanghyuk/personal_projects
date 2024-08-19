package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMarketHandler(t *testing.T) {

	app := fiber.New()

	m := NewMockMaketRetriever(t)
	setMarketRetrieverMock(m)

	f := MarketHandler{
		r: m,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("Market", func(t *testing.T) {
		reqBody := GetMarketParam{
			Date: "2024-08-20",
		}
		err := sendReqeust(app, "/market", "GET", reqBody)
		assert.NoError(t, err)
	})

	app.Shutdown()
}

func setMarketRetrieverMock(m *MockMaketRetriever) {
	m.On("RetrieveMarketSituation", mock.AnythingOfType("string")).Return("RetrieveMarketSituation Called", nil)
}
