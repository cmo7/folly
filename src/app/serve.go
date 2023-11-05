package app

import (
	"folly/src/app/routes"
	"folly/src/database"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Serve() {

	database.Connect()

	app := fiber.New(fiber.Config{
		AppName: viper.GetString("general.app.name"),
	})

	app.Mount("/api/v1", routes.ApiRouterV1())

	app.Listen(":" + viper.GetString("general.app.port"))
}
