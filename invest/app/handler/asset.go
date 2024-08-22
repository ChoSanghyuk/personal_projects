package handler

import (
	"fmt"
	m "invest/model"

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
	// router.Get("/:id/amount", h.AssetAmount)
	router.Get("/:id/hist", h.AssetHist)
	router.Post("/", h.SaveAssets)
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

// func (h *AssetHandler) AssetAmount(c *fiber.Ctx) error {

// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
// 	}

// 	fund, err := h.r.RetrieveAssetAmount(uint(id))
// 	if err != nil {
// 		return fmt.Errorf("RetrieveAssetAmount 오류 발생. %w", err)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fund)

// }

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

	var param m.SaveAssetParam
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}
	err = validCheck(param) // 포인터로 들어가도 validation 체크 되는지 확인
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}
	err = h.w.SaveAssetInfo(param.Name, param.Division, 0, "", param.Peak, 100, param.Bottom) // TODO ASSET 정보 계산
	if err != nil {
		return fmt.Errorf("SaveAssetInfo 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).SendString("자산 정보 저장 성공")
}
