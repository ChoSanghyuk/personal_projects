package handler

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// router.Get("/", h.TotalStatus)
// router.Post("/", h.AddFund)
// router.Get("/:id/hist", h.FundHist)
// router.Get("/:id/assets", h.FundAssets)

func TestFundHandler(t *testing.T) {

	app := fiber.New()

	readerMock := FundRetrieverMock{}
	writerMock := FundWriterMock{}

	f := FundHandler{
		r: readerMock,
		w: writerMock,
	}
	f.InitRoute(app)
	go func() {
		app.Listen(":3000")
	}()

	t.Run("전체 자금별 총액", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/funds", "GET", nil)
			assert.NoError(t, err)
		})

	})

	t.Run("자금 추가", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			param := AddFundReq{
				Name: "신규 자금",
			}
			err := sendReqeust(app, "/funds", "POST", param)
			assert.NoError(t, err)
		})

	})

	t.Run("자금 투자 이력 조회", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/funds/1/hist", "GET", nil)
			assert.NoError(t, err)
		})

	})

	t.Run("자금별 투자 종목 총액 조회", func(t *testing.T) {
		t.Run("성공 테스트", func(t *testing.T) {
			err := sendReqeust(app, "/funds/1/assets", "GET", nil)
			assert.NoError(t, err)
		})

	})

	app.Shutdown()
}
