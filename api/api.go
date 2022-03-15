package api

import "github.com/gofiber/fiber/v2"

func Start() error {
	app := fiber.New()

	app.Get("/api/xp/:player", xpOfPlayer)

	return app.Listen(":3000")
}
