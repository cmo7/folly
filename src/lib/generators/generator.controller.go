package generators

import (
	"bytes"
	"text/template"

	"github.com/gertd/go-pluralize"
)

func GenerateController(modelName string) string {

	pluralize := pluralize.NewClient()

	singular := pluralize.Singular(modelName)
	plural := pluralize.Plural(modelName)

	controllerTemplate := template.New("controller")

	controllerTemplate, err := controllerTemplate.Parse(controllerTemplateString)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}

	controllerTemplate.Execute(buf, map[string]string{
		"ModelName": modelName,
		"Singular":  singular,
		"Plural":    plural,
	})

	return buf.String()
}

var controllerTemplateString = `package controllers

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

func init() {
	RegisterController(generics.NewController[*models.{{.ModelName}}, *models.{{.ModelName}}DTO](generics.ResourceNames{
		Singular: "{{.Singular}}",
		Plural:   "{{.Plural}}",
	}))
}

`
