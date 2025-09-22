package main

import (
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/invest_indicator", func(c *fiber.Ctx) error {
		cmd := exec.Command("./auto_invest_indicator.sh")
		if err := cmd.Run(); err != nil {
			return c.SendString(err.Error())
		}
		return nil
	})
	app.Post("/lolche_bot", func(c *fiber.Ctx) error {
		cmd := exec.Command("./auto_lolche_bot.sh")
		if err := cmd.Run(); err != nil {
			return c.SendString(err.Error())
		}
		return nil
	})
	app.Post("/lolche_bot_rust", func(c *fiber.Ctx) error {
		cmd := exec.Command("./auto_lolche_bot_rust.sh")
		if err := cmd.Run(); err != nil {
			return c.SendString(err.Error())
		}
		return nil
	})
	app.Listen(":50000")
}
