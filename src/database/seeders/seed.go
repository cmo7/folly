package seeders

import (
	"folly/src/app/models"
	"folly/src/database/factories"
	"folly/src/lib/generics"
)

func Seed() error {

	// Create author role
	authorRole, err := factories.RoleFactory.CreateOneWith(generics.AttributeMap{
		"Name": "author",
	})
	if err != nil {
		return err
	}

	roleSlice := []models.Role{*authorRole}

	// Create 100 authors
	authors, err := factories.UserFactory.CreateManyWith(100, generics.AttributeMap{
		"Roles": roleSlice,
	})
	if err != nil {
		return err
	}

	// Create 100 posts for each author
	for _, author := range authors {
		derefAutor := *author
		_, err := factories.PostFactory.CreateManyWith(100, generics.AttributeMap{
			"Author": derefAutor,
		})
		if err != nil {
			return err
		}

	}

	return nil
}
