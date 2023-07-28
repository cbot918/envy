package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitApi() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	app, err := InitRoute(app)
	if err != nil {
		panic(err)
	}

	return app
}
