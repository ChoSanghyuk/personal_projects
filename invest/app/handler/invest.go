package handler

import "github.com/gofiber/fiber/v2"

type InvestHandler struct {
	r InvestRetriever
	w InvestSaver
}

func (h *InvestHandler) InitRoute(app *fiber.App) {

	router := app.Group("/invest")

	router.Get("/", nil)
	router.Post("/", nil)
}
