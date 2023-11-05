package controllers

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

func init() {
	RegisterController(generics.NewController[*models.User, *models.UserDTO](generics.ResourceNames{
		Singular: "user",
		Plural:   "users",
	}))
	RegisterController(generics.NewController[*models.Post, *models.PostDTO](generics.ResourceNames{
		Singular: "post",
		Plural:   "posts",
	}))
	RegisterController(generics.NewController[*models.Role, *models.RoleDTO](generics.ResourceNames{
		Singular: "role",
		Plural:   "roles",
	}))
}

var controllers = map[string]generics.GenericController{}

func RegisterController(controller generics.GenericController) {
	controllers[controller.GetResourceNames().Plural] = controller
}

func GetControllers() map[string]generics.GenericController {
	return controllers
}
