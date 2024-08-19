package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type InvestHandler struct {
	r InvestRetriever
	w InvestSaver
}

func (h *InvestHandler) InitRoute(app *fiber.App) {

	router := app.Group("/invest")

	router.Get("/hist", h.InvestHist)
	router.Post("/", h.SaveInvest)
}

func (h *InvestHandler) InvestHist(c *fiber.Ctx) error {

	var param GetInvestHistParam
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(&param)
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	investHist, err := h.r.RetrieveInvestHist(param.FundId, param.AssetId, param.StartDate, param.EndDate)
	if err != nil {
		return fmt.Errorf("RetrieveMarketSituation 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(investHist)
}

func (h *InvestHandler) SaveInvest(c *fiber.Ctx) error {

	param := SaveInvestParam{}
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	validCheck(&param)

	err = h.w.SaveInvest(param.FundId, param.AssetId, param.Price, param.Currency, param.Count)
	if err != nil {
		return fmt.Errorf("RetrieveMarketSituation 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).SendString("Invest 이력 저장")
}
