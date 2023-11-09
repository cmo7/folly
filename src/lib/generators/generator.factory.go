package generators

import (
	"bytes"
	"text/template"
)

func GenerateFactory(modelName string) string {
	factoryTemplate := template.New("factory")

	factoryTemplate, err := factoryTemplate.Parse(factoryTemplateString)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}

	factoryTemplate.Execute(buf, map[string]string{
		"ModelName": modelName,
	})

	return buf.String()
}

var factoryTemplateString = `package factories

import (
	"folly/src/app/models"
	"folly/src/lib/generics"
)

var {{.ModelName}}Factory = generics.NewFactory[*models.{{.ModelName}}](generics.GeneratorMap{
	//TODO: Add generators for each field of the model, e.g.:
	//"name": func() interface{} { return faker.Word() },
})

`
