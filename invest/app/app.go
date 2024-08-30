package app

import (
	"invest/app/handler"
	"invest/db"
	"invest/scrape"

	"github.com/gofiber/fiber/v2"
)

func Run(stg *db.Storage, scraper *scrape.Scraper) {

	app := fiber.New()

	handler.NewAssetHandler(stg, stg).InitRoute(app)
	handler.NewFundHandler(stg, stg, nil).InitRoute(app)
	handler.NewInvestHandler(stg).InitRoute(app)
	handler.NewMarketHandler(stg, stg).InitRoute(app)

	app.Listen(":3000")

}
