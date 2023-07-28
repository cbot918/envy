package main

import "github.com/gofiber/fiber/v2"

const (
	HOST = "127.0.0.1:3500"
)

func main() {
	app := fiber.New()

	app.Static("/", "./dist")

	app.Listen(HOST)
}
