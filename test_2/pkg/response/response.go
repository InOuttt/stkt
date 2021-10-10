package response

import "github.com/gofiber/fiber/v2"

func Generate(c *fiber.Ctx, code int, jsonData string) *fiber.Ctx {
	c.SendStatus(code)
	c.Response().Header.Add("content-type", "application/json")
	c.SendString(jsonData)

	return c
}
