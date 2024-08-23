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

	router.Get("/assets", h.TotalFundAssets)
	router.Get("/:id/assets", h.FundAsset)
	router.Get("/", h.TotalFunds)
	router.Get("/:id", h.Fund) // routing 포함 위치 중요. /assets보다 아래에 있어야 함.
}

func (h *FundHandler) TotalFunds(c *fiber.Ctx) error {

	funds, err := h.r.RetrieveFundAmount()
	if err != nil {
		return fmt.Errorf("retrieveFundAmount시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

func (h *FundHandler) Fund(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetrieveFundHistById(uint(id))
	if err != nil {
		return fmt.Errorf("retrieveFundAmount시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)
}

func (h *FundHandler) TotalFundAssets(c *fiber.Ctx) error {

	funds, err := h.r.RetreiveInvestHistOfFund()
	if err != nil {
		return fmt.Errorf("retreiveAssetOfFund 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

func (h *FundHandler) FundAsset(c *fiber.Ctx) error {

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
