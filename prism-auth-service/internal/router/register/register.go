package register

import (
	"context"
	"prism-auth-service/internal/database"
	"prism-auth-service/pkg/email"
	passwordhash "prism-auth-service/pkg/passwordHash"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	mail, password := c.FormValue("e"), c.FormValue("p")

	if !email.Validate(mail) {
		return c.JSON(fiber.Map{
			"Status": "failed",
			"msg":    "Invalid E-Mail address",
		})
	}

	var count int
	database.Get().QueryRow(context.Background(), "SELECT COUNT(*) FROM users WHERE email=$1", mail).Scan(&count)

	if count >= 1 {
		return c.JSON(fiber.Map{
			"Status": "failed",
			"msg":    "User exists",
		})
	}

	passwordHashed, _ := passwordhash.New(password)

	database.Get().Query(
		context.Background(),
		"INSERT INTO users(email, password_hash, role, created_at) VALUES($1, $2, $3, $4)", mail, passwordHashed, "default", time.Now(),
	)

	return c.SendString("test")
}
