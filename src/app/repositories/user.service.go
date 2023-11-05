package repositories

import (
	"folly/src/app/models"
	"folly/src/database"
	"folly/src/lib/generics"
)

type UserRepository struct {
	generics.GenericRepository[*models.User, *models.UserDTO]
	// Add custom method signatures here
	FindByEmail  func(email string, relations []string) (models.User, error)
	IsEmailTaken func(email string) bool
}

var UserRepositoryGORM UserRepository = UserRepository{
	// Add custom method imlementations here
	FindByEmail: func(email string, relations []string) (models.User, error) {
		var user models.User
		result := database.DB.
			Scopes(generics.Preload(relations)).
			Where("email = ?", email).
			First(&user)
		return user, result.Error
	},
	IsEmailTaken: func(email string) bool {
		var count int64
		database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
		return count > 0
	},
}
