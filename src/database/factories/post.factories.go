package factories

import (
	"folly/src/app/models"
	"folly/src/app/repositories"

	"folly/src/lib/generics"
)

var PostFactory = generics.NewFactory[*models.Post](generics.GeneratorMap{
	"Title": func() interface{} {
		return faker.Sentence(5)
	},
	"Content": func() interface{} {
		return faker.Paragraph(5, 6, 5, "\n")
	},
	"Author": func() interface{} {
		author, err := repositories.UserRepositoryGORM.FindOneRandom()
		if err != nil {
			panic(err)
		}
		return author
	},
})
