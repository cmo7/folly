package generics

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewGenericRouter(controller GenericController) *fiber.App {
	app := fiber.New()

	app.Get("/", controller.GetAll()).Name(fmt.Sprintf("GetAll%s", controller.GetResourceNames().Plural))
	app.Get("/:id", controller.Get()).Name(fmt.Sprintf("Get%s", controller.GetResourceNames().Singular))
	app.Post("/", controller.Create()).Name(fmt.Sprintf("Create%s", controller.GetResourceNames().Singular))
	app.Put("/:id", controller.Update()).Name(fmt.Sprintf("Update%s", controller.GetResourceNames().Singular))
	app.Delete("/:id", controller.Delete()).Name(fmt.Sprintf("Delete%s", controller.GetResourceNames().Singular))

	return app
}
