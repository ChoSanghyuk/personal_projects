package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MarketHandler struct {
	r MaketRetriever
}

func (h *MarketHandler) InitRoute(app *fiber.App) {

	router := app.Group("/market")

	router.Get("/", h.Market)

}

func (h *MarketHandler) Market(c *fiber.Ctx) error {

	param := GetMarketParam{}
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	validCheck(&param)

	assets, err := h.r.RetrieveMarketSituation(param.Date)
	if err != nil {
		return fmt.Errorf("RetrieveMarketSituation 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(assets)
}