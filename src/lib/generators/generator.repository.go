package generators

import (
	"bytes"
	"text/template"
)

func GenerateRepository(modelName string) string {
	repositoryTemplate := template.New("repository")

	repositoryTemplate, err := repositoryTemplate.Parse(repositoryTemplateString)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}

	repositoryTemplate.Execute(buf, map[string]string{
		"ModelName": modelName,
	})

	return buf.String()
}

var repositoryTemplateString = `package repositories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

type {{.ModelName}}Repository struct {
	generics.GenericRepository[*models.{{.ModelName}}, *models.{{.ModelName}}DTO]
	// TODO: Add custom method signatures here
}

var {{.ModelName}}RepositoryGORM {{.ModelName}}Repository = {{.ModelName}}Repository{
	// TODO: Add custom method imlementations here
}
`
