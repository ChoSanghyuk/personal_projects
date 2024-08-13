package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestFundHandler(t *testing.T) {

	app := fiber.New()
	f := FundHandler{
		r: FundRetrieverMock{},
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("TotalFunds", func(t *testing.T) {
		sendReqeust(t, app, "/funds")
	})

	t.Run("Fund", func(t *testing.T) {
		sendReqeust(t, app, "/funds/1")
	})

	t.Run("TotalFundAssets", func(t *testing.T) {
		sendReqeust(t, app, "/funds/assets")
	})

	t.Run("FundAsset", func(t *testing.T) {
		sendReqeust(t, app, "/funds/1/assets")
	})

	app.Shutdown()
}

/*
************************************************
Inner Function
*************************************************
*/
func sendReqeust(t *testing.T, app *fiber.App, url string) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Error("Error Occurred", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Error("Response status should be 200")
	}

	body, _ := io.ReadAll(resp.Body)
	t.Log(string(body))
}

/*
************************************************
Mock
*************************************************
*/
type FundRetrieverMock struct {
}

func (m FundRetrieverMock) RetrieveFundAmount() (any, error) {
	return "hello1", nil
}
func (m FundRetrieverMock) RetrieveFundAmountById(id uint) (any, error) {
	return "hello2", nil

}
func (m FundRetrieverMock) RetreiveAssetOfFund() (any, error) {
	return "hello3", nil

}
func (m FundRetrieverMock) RetreiveAssetOfFundById(id uint) (any, error) {
	return "hello4", nil

}
