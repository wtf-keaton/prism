package token

import "github.com/gofiber/fiber/v2"

func ValidateToken(c *fiber.Ctx) error {
	return c.SendString("test")
}

func Refresh(c *fiber.Ctx) error {
	return c.SendString("test2")
}
