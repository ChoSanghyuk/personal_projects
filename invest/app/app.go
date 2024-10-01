package app

import (
	"fmt"
	"invest/app/handler"
	"invest/db"
	"invest/scrape"

	"github.com/gofiber/fiber/v2"
)

func Run(stg *db.Storage, scraper *scrape.Scraper) {

	app := fiber.New()

	handler.NewAssetHandler(stg, stg, scraper).InitRoute(app)
	handler.NewFundHandler(stg, stg, scraper).InitRoute(app)
	handler.NewInvestHandler(stg, stg).InitRoute(app)
	handler.NewMarketHandler(stg, stg).InitRoute(app)

	app.Get("/shutdown", func(c *fiber.Ctx) error {

		fmt.Println("Shutting Down")
		panic("SHUTDOWN")
		// go func() {
		// 	if err := app.Shutdown(); err != nil {
		// 		panic("SHUTDOWN ERROR")
		// 	}
		// }()
		return c.SendString("Shutting Down")
	})

	app.Listen(":3000")

}
