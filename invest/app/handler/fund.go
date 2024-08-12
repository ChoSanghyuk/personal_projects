package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FundHandler struct {
	db FundRetriever
}

func (f *FundHandler) InitRoute(app *fiber.App) {

	router := app.Group("/fund")
	router.Get("", f.TotalFunds)
	router.Get("/:id", f.Fund)
	router.Get("/assets", f.TotalFundAssets)
	router.Get("/assets/:id", f.FundAsset)
}

func (f *FundHandler) TotalFunds(c *fiber.Ctx) error {

	funds, err := f.db.RetrieveFundAmount()
	if err != nil {
		return fmt.Errorf("retrieveFundAmount시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

func (f *FundHandler) Fund(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := f.db.RetrieveFundAmountById(uint(id))
	if err != nil {
		return fmt.Errorf("retrieveFundAmount시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)
}

func (f *FundHandler) TotalFundAssets(c *fiber.Ctx) error {

	funds, err := f.db.RetreiveAssetOfFund()
	if err != nil {
		return fmt.Errorf("retreiveAssetOfFund 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(funds)
}

func (f *FundHandler) FundAsset(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("파라미터 id 조회 시 오류 발생. %w", err)
	}

	fund, err := f.db.RetreiveAssetOfFundById(uint(id))
	if err != nil {
		return fmt.Errorf("retreiveAssetOfFundById 시 오류 발생. %w", err)
	}

	return c.Status(fiber.StatusOK).JSON(fund)
}
