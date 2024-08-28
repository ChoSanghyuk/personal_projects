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
)

func TestAssetGetHandler(t *testing.T) {

	app := fiber.New()

	readerMock := AssetRetrieverMock{}

	f := AssetHandler{
		r: readerMock,
		w: nil,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("Get Asset", func(t *testing.T) {
		err := sendReqeust(app, "/assets/1", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("Get Asset List", func(t *testing.T) {
		err := sendReqeust(app, "/assets/list", "GET", nil)
		assert.NoError(t, err)
	})

	t.Run("Get wrong url", func(t *testing.T) {
		err := sendReqeust(app, "/assets/true", "GET", nil)
		assert.Error(t, err)
	})

	app.Shutdown()
}

func TestAssetPostHandler(t *testing.T) {

	app := fiber.New()

	writerMock := AssetInfoSaverMock{}

	f := AssetHandler{
		r: nil,
		w: writerMock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("SaveAssets", func(t *testing.T) {
		reqBody := AddAssetReq{
			ID:       1,
			Name:     "sample",
			Category: 1,
			Currency: "USD",
			Top:      500,
			Bottom:   400,
		}
		err := sendReqeust(app, "/assets", "POST", reqBody)
		assert.NoError(t, err)
	})

	t.Run("SaveAssets_InvalidReq", func(t *testing.T) {
		reqBody := AddAssetReq{}
		err := sendReqeust(app, "/assets", "POST", reqBody)
		assert.Error(t, err)
	})
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
