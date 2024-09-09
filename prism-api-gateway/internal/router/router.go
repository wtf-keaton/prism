package router

import (
	"prism-api-gateway/internal/router/handlers"

	"github.com/gofiber/fiber/v2"
)

func Serve() (app *fiber.App) {
	app = fiber.New()

	api := app.Group("/api/v1")
	api.Post("/sign_in", handlers.SignIn)
	api.Post("/sign_up", handlers.SignUp)

	return
}
