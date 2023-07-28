package internal

import "github.com/gofiber/fiber/v2"

func InitRoute(app *fiber.App) (*fiber.App, error) {

	c := NewController()

	app.Post("/install", c.Install)

	return app, nil
}
