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

// router.Post("/", h.AddAsset)
// router.Put("/", h.UpdateAsset)
// router.Delete("/", h.DeleteAsset)

func TestAssetGetHandler(t *testing.T) {

	app := fiber.New()

	readerMock := AssetRetrieverMock{}
	writerMock := AssetInfoSaverMock{}

	f := AssetHandler{
		r: readerMock,
		w: writerMock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("종목 리스트 조회 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/assets/list", "GET", nil)
			assert.NoError(t, err)
		})
	})

	t.Run("종목 정보 조회 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/assets/1", "GET", nil)
			assert.NoError(t, err)
		})
	})

	t.Run("종목 투자 이력 조회 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/assets/1/hist", "GET", nil)
			assert.NoError(t, err)
		})
	})

	t.Run("종목 추가 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			param := AddAssetReq{
				Name:      "종목",
				Category:  5,
				Currency:  "WON",
				Top:       500,
				Bottom:    400,
				SellPrice: 480,
				BuyPrice:  450,
				Path:      "",
			}
			err := sendReqeust(app, "/assets/", "POST", param)
			assert.NoError(t, err)
		})

		t.Run("실패 테스트 - 필수 파라미터 미존재", func(t *testing.T) {
			param := AddAssetReq{
				// Name      : "종목",
				Category:  5,
				Currency:  "WON",
				Top:       500,
				Bottom:    400,
				SellPrice: 480,
				BuyPrice:  450,
				Path:      "",
			}
			err := sendReqeust(app, "/assets/", "POST", param)
			assert.NoError(t, err)
		})
	})

	t.Run("종목 갱신 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			param := UpdateAssetReq{
				ID:        1,
				Name:      "종목",
				Category:  5,
				Currency:  "WON",
				Top:       500,
				Bottom:    400,
				SellPrice: 480,
				BuyPrice:  450,
				Path:      "",
			}
			err := sendReqeust(app, "/assets/", "PUT", param)
			assert.NoError(t, err)
		})

		t.Run("실패 테스트 - 필수 파라미터 미존재", func(t *testing.T) {
			param := UpdateAssetReq{
				// ID:        1,
				Name:      "종목",
				Category:  5,
				Currency:  "WON",
				Top:       500,
				Bottom:    400,
				SellPrice: 480,
				BuyPrice:  450,
				Path:      "",
			}
			err := sendReqeust(app, "/assets/", "PUT", param)
			assert.NoError(t, err)
		})
	})

	t.Run("종목 삭제 테스트", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			param := DeleteAssetReq{
				ID: 1,
			}
			err := sendReqeust(app, "/assets/", "DELETE", param)
			assert.NoError(t, err)
		})

		t.Run("실패 테스트 - 필수 파라미터 미존재", func(t *testing.T) {
			param := DeleteAssetReq{
				// ID: 1,
			}
			err := sendReqeust(app, "/assets/", "DELETE", param)
			assert.NoError(t, err)
		})
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
