package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func FundHandlerTest(t *testing.T) {

	app := fiber.New()
	f := FundHandler{
		db: FundRetrieverMock{},
	}
	f.InitRoute(app)

	t.Run("/fund", func(t *testing.T) {

	})
}

type FundRetrieverMock struct {
}

func (m FundRetrieverMock) RetrieveFundAmount() (any, error) {
	return "hello", nil
}
func (m FundRetrieverMock) RetrieveFundAmountById(id uint) (any, error) {
	return "hello", nil

}
func (m FundRetrieverMock) RetreiveAssetOfFund() (any, error) {
	return "hello", nil

}
func (m FundRetrieverMock) RetreiveAssetOfFundById(id uint) (any, error) {
	return "hello", nil

}
