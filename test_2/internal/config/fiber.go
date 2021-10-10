package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inouttt/stkt/test_2/internal/constant"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			c.Context().Error(constant.GetError(constant.CodeBadRequest))
			return nil
		},
	}
}
