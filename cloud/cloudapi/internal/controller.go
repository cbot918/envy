package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Controller struct {
	Log zerolog.Logger
}

func NewController(log zerolog.Logger) *Controller {
	return &Controller{
		Log: log,
	}
}

func (ctrl *Controller) Login(c *fiber.Ctx) error {
	// req := InstallRequest{}
	// if err := c.BodyParser(&req); err != nil {
	// 	return err
	// }
	// Lgj(req)

	// return c.JSON(map[string]string{
	// 	"message": "receive post",
	// })

	ctrl.Log.Info().Msg("login request")

	return c.JSON(map[string]string{
		"cloudapi:": "hihi",
	})
}
