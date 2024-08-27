package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	model "invest/model"
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
Inner Function
*************************************************
*/
func sendReqeust(app *fiber.App, url string, method string, reqBody any) error {

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("Error Occurred %w", err)
	}

	var req *http.Request
	switch method {
	case "POST":
		req, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	default:
		req, _ = http.NewRequest(http.MethodGet, url, bytes.NewBuffer(bodyBytes))
	}
	req.Header.Set("Content-Type", "application/json") // 중요!!. 생략 시, 파싱 오류 발생
	resp, err := app.Test(req, -1)
	if err != nil {
		return fmt.Errorf("Error Occurred %w", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		return fmt.Errorf("Response status should be 200. Status: %d", resp.StatusCode)
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

	m.On("RetreiveAssetOfFundById", mock.AnythingOfType("uint")).Return(&model.Fund{ID: 3, Name: "개인"}, nil)
	m.On("RetreiveInvestHistOfFund").Return([]model.Fund{}, nil)
	m.On("RetrieveFundAmount").Return([]model.Fund{{ID: 3, Name: "개인"}}, nil)
	m.On("RetreiveInvestHistOfFundById", mock.AnythingOfType("uint")).Return(&model.Fund{}, nil)

	return nil
}
