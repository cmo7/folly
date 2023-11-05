package factories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

var RoleFactory = generics.NewFactory[*models.Role](generics.GeneratorMap{
	"name": func() interface{} { return faker.Word() },
})
