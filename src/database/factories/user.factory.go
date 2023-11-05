package factories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

var UserFactory = generics.NewFactory[*models.User](generics.GeneratorMap{
	"Email":     func() interface{} { return faker.Email() },
	"Password":  func() interface{} { return faker.Password(true, true, true, true, false, 10) },
	"FirstName": func() interface{} { return faker.FirstName() },
	"LastName":  func() interface{} { return faker.LastName() },
})
