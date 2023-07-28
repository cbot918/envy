package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Route struct {
	Log zerolog.Logger
}

func NewRoute(log zerolog.Logger) *Route {
	return &Route{
		Log: log,
	}
}

func (r *Route) InitRoute(app *fiber.App) (*fiber.App, error) {

	c := NewController(r.Log)

	app.Get("/login", c.Login)

	return app, nil
}
