package routes

import (
	"folly/src/app/controllers"
	v1 "folly/src/app/routes/v1"
	"folly/src/lib/generics"

	"github.com/gofiber/fiber/v2"
)

func ApiRouterV1() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Mount("/auth", v1.AuthRoutes())

	for key, controller := range controllers.GetControllers() {
		app.Mount("/"+key, generics.NewGenericRouter(controller))
	}

	app.Get("/info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"routes": app.GetRoutes(),
		})
	})

	return app
}
