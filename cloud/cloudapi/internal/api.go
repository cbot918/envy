package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
)

type CloudAPI struct {
	Cfg    *Config
	Log    zerolog.Logger
	Server *fiber.App
}

func NewCloudAPI(cfg *Config, log zerolog.Logger) (*CloudAPI, error) {
	api := new(CloudAPI)
	api.Cfg = cfg
	api.Log = log

	api.Server = fiber.New()
	api.Server.Use(cors.New())

	var err error
	api.Server, err = NewRoute(log).InitRoute(api.Server)
	if err != nil {
		return nil, err
	}

	return api, nil
}
