package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

/*
************************************************
Inner Function
*************************************************
*/
func sendReqeust(app *fiber.App, url string, reqBody any) error {

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("Error Occurred %w", err)
	}
	req, _ := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(bodyBytes))

	resp, err := app.Test(req, -1)
	if err != nil {
		return fmt.Errorf("Error Occurred %w", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("Response status should be 200")
	}

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	return nil
}

/*
************************************************
Mock
*************************************************
*/

func setFundRetrieverMock(m *MockFundRetriever) error {

	m.On("RetrieveFundAmount").Return("hello1", nil)
	m.On("RetrieveFundAmountById", mock.AnythingOfType("uint")).Return("hello2", nil)
	m.On("RetreiveAssetOfFund").Return("hello3", nil)
	m.On("RetreiveAssetOfFundById", mock.AnythingOfType("uint")).Return("hello4", nil)

	return nil
}

/*
************************************************
Mock_OLd
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
