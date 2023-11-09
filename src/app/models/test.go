package models

import (
	"folly/src/database"
	"folly/src/lib/common"
)

func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &Aguacate{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}

type Aguacate struct {
	common.CommonEntity `gorm:"embedded"`
	// TODO: Add fields here. Dont forget to add the gorm tags
}

type AguacateDTO struct {
	common.CommonDTO `json:",inline,omitempty"`
	// TODO: Add fields here. Dont forget to add the json tags
}

func (u Aguacate) ToDto() common.DTO {
	// TODO: Implement this method

	dto := &AguacateDTO{}
	// TODO: Convert the fields from the entity to the dto
	return dto
}

func (u AguacateDTO) ToEntity() common.Entity {
	// TODO: Implement this method
	entity := &Aguacate{}
	// TODO: Convert the fields from the dto to the entity
	return entity
}

