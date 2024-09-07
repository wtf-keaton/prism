package login

import (
	"prism-auth-service/pkg/email"

	"github.com/gofiber/fiber/v2"
)

func SignIn(c *fiber.Ctx) error {
	mail, _ := c.FormValue("e"), c.FormValue("p")

	if !email.Validate(mail) {
		return c.JSON(fiber.Map{
			"Status": "failed",
			"msg":    "Invalid E-Mail address",
		})
	}
	return c.SendString("test")
}
