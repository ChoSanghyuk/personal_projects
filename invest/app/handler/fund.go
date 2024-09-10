package handler

import (
	"fmt"
	"invest/model"

	"github.com/gofiber/fiber/v2"
)

type FundHandler struct {
	r FundRetriever
	w FundWriter
	e ExchageRateGetter
}

func NewFundHandler(r FundRetriever, w FundWriter, e ExchageRateGetter) *FundHandler {
	return &FundHandler{
		r: r,
		w: w,
		e: e,
	}
}

func (h *FundHandler) InitRoute(app *fiber.App) {
	router := app.Group("/funds")

	router.Get("/", h.TotalStatus)
	router.Post("/", h.AddFund)
	router.Get("/:id/hist", h.FundHist)
	router.Get("/:id/assets", h.FundAssets)
}

// 총 자금 금액
func (h *FundHandler) TotalStatus(c *fiber.Ctx) error {

	var exchangeRate float64 = h.e.ExchageRate()

	summarys, err := h.r.RetreiveFundsSummaryOrderByFundId()
	if err != nil {
		return fmt.Errorf("RetreiveFundSummary 오류 발생. %w", err)
	}

	funds := make([]*TotalStatusResp, 0)
	for _, s := range summarys {

		id := int(s.FundID)
		if id > len(funds) {
			funds = append(funds, &TotalStatusResp{})
		}

		fund := funds[id-1]

		if fund.ID == 0 {
			fund.ID = s.FundID
			fund.Name = s.Fund.Name
		}

		if s.Asset.Currency == model.KRW.String() {
			fund.Amount = exchangeRate + s.Sum*exchangeRate
		} else {
			fund.Amount = fund.Amount + s.Sum
		}
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

func (h *FundHandler) AddFund(c *fiber.Ctx) error {

	var param AddFundReq
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(param) // 포인터로 들어가도 validation 체크 되는지 확인
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	err = h.w.SaveFund(param.Name)
	if err != nil {
		return fmt.Errorf("SaveFund 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).SendString("자금 정보 저장 성공")
}

// 자금별 보유 자산
func (h *FundHandler) FundAssets(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	funds, err := h.r.RetreiveFundSummaryByFundId(uint(id))
	if err != nil {
		return fmt.Errorf("RetreiveFundSummaryById 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

// 자금별 투자 이력
func (h *FundHandler) FundHist(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetreiveAFundInvestsById(uint(id))
	if err != nil {
		return fmt.Errorf("RetreiveAFundInvestsById 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)
}
