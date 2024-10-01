package handler

import (
	"errors"
	"fmt"
	"invest/model"

	"github.com/gofiber/fiber/v2"
)

type InvestHandler struct {
	r  AssetRetriever
	w  InvestSaver
	cm map[model.Currency]uint
}

func (h *InvestHandler) InitRoute(app *fiber.App) {
	router := app.Group("/invest")
	router.Post("/", h.SaveInvest)
}

func NewInvestHandler(r AssetRetriever, w InvestSaver) *InvestHandler {

	cm := make(map[model.Currency]uint)
	li, err := r.RetrieveAssetList()
	if err != nil {
		panic("InvestHandler 기동시 오류. Shutdown")
	}

	for _, a := range li {
		if a.Name == model.KRW.String() {
			cm[model.KRW] = a.ID
		} else if a.Name == model.USD.String() {
			cm[model.USD] = a.ID
		}
	}

	return &InvestHandler{
		r:  r,
		w:  w,
		cm: cm,
	}
}

func (h *InvestHandler) SaveInvest(c *fiber.Ctx) error {

	param := SaveInvestParam{}
	err := c.BodyParser(&param)
	if err != nil {
		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
	}

	err = validCheck(&param)
	if err != nil {
		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
	}

	var assetId uint

	// assetId 미존재 시, name 혹은 code로 Id 구해옴
	if param.AssetId != 0 {
		assetId = param.AssetId
	} else if param.AssetName != "" {
		assetId = h.r.RetrieveAssetIdByName(param.AssetName)
	} else if param.AssetCode != "" {
		assetId = h.r.RetrieveAssetIdByCode(param.AssetName)
	}
	if assetId == 0 {
		return errors.New("parameter asset 정보 없음")
	}

	// 투자 이력 저장
	err = h.w.SaveInvest(param.FundId, assetId, param.Price, param.Count)
	if err != nil {
		return fmt.Errorf("SaveInvest 오류 발생. %w", err)
	}

	// 투자 요약 갱신
	err = h.w.UpdateInvestSummary(param.FundId, assetId, param.Count, param.Price)
	if err != nil {
		return fmt.Errorf("UpdateInvestSummaryCount 오류 발생. %w", err)
	}

	// 현금/달러 갱신
	asset, err := h.r.RetrieveAsset(assetId)
	if err != nil {
		return fmt.Errorf("RetrieveAsset 오류 발생. %w", err)
	}

	if asset.Currency == model.KRW.String() && asset.Name != model.KRW.String() {
		err = h.w.UpdateInvestSummary(param.FundId, h.cm[model.KRW], -1*param.Price*param.Count, 1)
	} else if asset.Currency == model.USD.String() && asset.Name != model.USD.String() {
		err = h.w.UpdateInvestSummary(param.FundId, h.cm[model.KRW], -1*param.Price*param.Count, 1)
	}
	if err != nil {
		return fmt.Errorf("UpdateInvestSummaryCount 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).SendString("Invest 이력 저장 성공")
}

// func (h *InvestHandler) InvestHist(c *fiber.Ctx) error {

// 	var param model.GetInvestHistParam
// 	err := c.BodyParser(&param)
// 	if err != nil {
// 		return fmt.Errorf("파라미터 BodyParse 시 오류 발생. %w", err)
// 	}

// 	err = validCheck(&param)
// 	if err != nil {
// 		return fmt.Errorf("파라미터 유효성 검사 시 오류 발생. %w", err)
// 	}

// 	investHist, err := h.r.RetrieveInvestHist(param.FundId, param.AssetId, param.StartDate, param.EndDate)
// 	if err != nil {
// 		return fmt.Errorf("RetrieveMarketSituation 오류 발생. %w", err)
// 	}

// 	return c.Status(fiber.StatusOK).JSON(investHist)
// }
