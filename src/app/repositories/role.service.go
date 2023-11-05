package repositories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

type RoleRepository struct {
	generics.GenericRepository[*models.Role, *models.RoleDTO]
}

var RoleRepositoryGORM RoleRepository = RoleRepository{}
