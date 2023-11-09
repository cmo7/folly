package generators

import (
	"bytes"
	"text/template"
)

func GenerateModel(modelName string) string {
	modelTemplate := template.New("model")

	modelTemplate, err := modelTemplate.Parse(modelTemplateString)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}

	modelTemplate.Execute(buf, map[string]string{
		"ModelName":    modelName,
		"GormEmbedded": "`gorm:\"embedded\"`",
		"JSONInline":   "`json:\",inline,omitempty\"`",
	})

	return buf.String()
}

var modelTemplateString = `package models

import (
	"folly/src/database"
	"folly/src/lib/common"
)

func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &{{.ModelName}}{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}

type {{.ModelName}} struct {
	common.CommonEntity {{.GormEmbedded}}
	// TODO: Add fields here. Dont forget to add the gorm tags
}

type {{.ModelName}}DTO struct {
	common.CommonDTO {{.JSONInline}}
	// TODO: Add fields here. Dont forget to add the json tags
}

func (u {{.ModelName}}) ToDto() common.DTO {
	// TODO: Implement this method

	dto := &{{.ModelName}}DTO{}
	// TODO: Convert the fields from the entity to the dto
	return dto
}

func (u {{.ModelName}}DTO) ToEntity() common.Entity {
	// TODO: Implement this method
	entity := &{{.ModelName}}{}
	// TODO: Convert the fields from the dto to the entity
	return entity
}

`
