package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AssetHandler struct {
	r AssetRetriever
	w AssetInfoSaver
}

func (h *AssetHandler) InitRoute(app *fiber.App) {

	router := app.Group("/assets")

	router.Get("/", h.Assets)
	router.Get("/:id", h.AssetInfo)
	router.Get("/:id/amount", h.AssetAmount)
	router.Get("/:id/hist", h.AssetHist)
}

func (h *AssetHandler) Assets(c *fiber.Ctx) error {
	assets, err := h.r.RetrieveAssetList()
	if err != nil {
		return fmt.Errorf("RetrieveAssetList 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(assets)
}

func (h *AssetHandler) AssetInfo(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetrieveAssetInfo(uint(id))
	if err != nil {
		return fmt.Errorf("RetrieveAssetInfo 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)

}

func (h *AssetHandler) AssetAmount(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetrieveAssetAmount(uint(id))
	if err != nil {
		return fmt.Errorf("RetrieveAssetAmount 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)

}

func (h *AssetHandler) AssetHist(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := h.r.RetrieveAssetHist(uint(id))
	if err != nil {
		return fmt.Errorf("RetrieveAssetHist 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)

}

func (h *AssetHandler) SaveAssets(c *fiber.Ctx) error {

	param := Asset{}
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = h.w.SaveAssetInfo(param)
	if err != nil {
		return fmt.Errorf("SaveAssetInfo 시 오류 발생. %w", err)
	}

	c.Status(fiber.StatusOK)
	return nil
}
