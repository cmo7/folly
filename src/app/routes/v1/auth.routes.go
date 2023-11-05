package v1

import (
	"folly/src/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes() *fiber.App {

	var router = fiber.New()
	controller := controllers.AuthController

	router.Post("/login", controller.Login)
	router.Post("/register", controller.Register)
	router.Post("/logout", controller.Logout)
	router.Post("/refresh-access-token", controller.RefreshAccessToken)

	return router
}
