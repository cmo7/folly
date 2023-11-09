package generators

import (
	"fmt"
	"os"

	"github.com/iancoleman/strcase"
)

func Generate(modelName string) {
	modelFormat := "src/app/models/%s.model.go"
	factoryFormat := "src/database/factories/%s.factory.go"
	controllerFormat := "src/app/controllers/%s.controller.go"
	repositoryFormat := "src/app/repositories/%s.repository.go"

	writeToFile(modelFormat, GenerateModel(modelName), modelName)
	writeToFile(factoryFormat, GenerateFactory(modelName), modelName)
	writeToFile(controllerFormat, GenerateController(modelName), modelName)
	writeToFile(repositoryFormat, GenerateRepository(modelName), modelName)
}

func writeToFile(format string, content string, modelName string) {
	fileName := strcase.ToDelimited(modelName, '.')
	path := fmt.Sprintf(format, fileName)
	os.WriteFile(path, []byte(content), 0644)
}
