package repositories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

type PostRepository struct {
	generics.GenericRepository[*models.Post, *models.PostDTO]
	// Add custom method signatures here
}

var PostRepositoryGorm PostRepository = PostRepository{
	// Add custom method imlementations here
}
