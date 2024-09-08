package router

import (
	"prism-frontend-service/internal/router/authentication"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Serve() (app *fiber.App) {
	viewEngine := html.New("./web/templates", ".html")
	viewEngine.Reload(true)

	app = fiber.New(fiber.Config{
		Prefork: true,
		Views:   viewEngine,
	})

	app.Static("/assets", "./web/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/sign_in")
	})

	app.Get("/sign_in", authentication.SignIn)
	app.Get("/sign_up", authentication.SignUp)

	app.Post("/actions/sign_in", authentication.LogIn)
	return
}
