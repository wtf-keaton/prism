package login

import "github.com/gofiber/fiber/v2"

func SignIn(c *fiber.Ctx) error {

	return c.SendString("test")
}
