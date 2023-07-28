package internal

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Install(c *fiber.Ctx) error {
	req := InstallRequest{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	Lgj(req)

	return c.JSON(map[string]string{
		"message": "receive post",
	})
}
