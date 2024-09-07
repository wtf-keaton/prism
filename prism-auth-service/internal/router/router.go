package router

import (
	"prism-auth-service/internal/router/login"
	"prism-auth-service/internal/router/register"
	"prism-auth-service/internal/router/token"

	"github.com/gofiber/fiber/v2"
)

func Serve() (app *fiber.App) {
	app = fiber.New()

	api := app.Group("/api")
	api.Post("/sign_in", login.SignIn)
	api.Post("/sign_up", register.SignUp)

	api.Post("/validate", token.ValidateToken)
	api.Post("/refresh", token.RefreshToken)

	return
}
