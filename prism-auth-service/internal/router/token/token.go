package token

import "github.com/gofiber/fiber/v2"

func ValidateToken(c *fiber.Ctx) error {
	return c.SendString("test")
}

func RefreshToken(c *fiber.Ctx) error {
	return c.SendString("test2")
}
