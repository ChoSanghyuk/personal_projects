package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MarketHandler struct {
	r MaketRetriever
	w MarketSaver
}

func (h *MarketHandler) InitRoute(app *fiber.App) {

	router := app.Group("/market")
	router.Get("/", h.Market)
	router.Get("/indicator", h.MarketIndicator)
	router.Post("/", h.ChangeMarketStatus)
}

func (h *MarketHandler) Market(c *fiber.Ctx) error {

	var param MarketStatusParam
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(&param)
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	assets, err := h.r.RetrieveMarketStatus(param.Date)
	if err != nil {
		return fmt.Errorf("RetrieveMarketStatus 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(assets)
}

func (h *MarketHandler) MarketIndicator(c *fiber.Ctx) error {

	var param MarketStatusParam
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(&param)
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	dailyIdx, cliIdx, err := h.r.RetrieveMarketIndicator(param.Date)
	if err != nil {
		return fmt.Errorf("RetrieveMarketIndicator 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON([]any{dailyIdx, cliIdx})
}

func (h *MarketHandler) ChangeMarketStatus(c *fiber.Ctx) error {

	var param SaveMarketStatusParam
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(&param)
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	err = h.w.SaveMarketStatus(param.Status)
	if err != nil {
		return fmt.Errorf("RetrieveMarketStatus 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).SendString("시장 상태 저장 성공")

}
