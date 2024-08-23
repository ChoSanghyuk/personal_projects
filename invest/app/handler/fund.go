package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FundHandler struct {
	r FundRetriever
}

func (h *FundHandler) InitRoute(app *fiber.App) {

	router := app.Group("/funds")

	router.Get("/", h.TotalFunds)
	router.Get("/:id/assets", h.TotalFundAssets)
	router.Get("/:id/hist", h.FundHist)

}

// 총 자금 금액
func (h *FundHandler) TotalFunds(c *fiber.Ctx) error {

	funds, err := h.r.RetreiveInvestHistOfFunds()
	if err != nil {
		return fmt.Errorf("RetreiveInvestHistOfFunds 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

// 자금별 보유 자산
func (h *FundHandler) TotalFundAssets(c *fiber.Ctx) error {

	funds, err := h.r.RetreiveInvestHistOfFunds()
	if err != nil {
		return fmt.Errorf("retreiveAssetOfFund 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

// 자금별 투자 이력
func (h *FundHandler) FundHist(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetreiveAssetOfFundById(uint(id))
	if err != nil {
		return fmt.Errorf("retreiveAssetOfFundById 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)
}
