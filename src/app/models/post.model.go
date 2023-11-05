package models

import (
	"folly/src/database"
	"folly/src/lib/common"

	"github.com/google/uuid"
)

func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &Post{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}

type Post struct {
	common.CommonEntity `gorm:"embedded"`
	Title               string `gorm:"not null"`
	Content             string `gorm:"not null"`
	AuthorID            uuid.UUID
	Author              User `gorm:"foreignKey:AuthorID"`
}

type PostDTO struct {
	common.CommonDTO `json:",inline,omitempty"`
	Title            string  `json:"title"`
	Content          string  `json:"content"`
	Author           UserDTO `json:"author,omitempty"`
}

func (p Post) ToDto() common.DTO {
	author := p.Author.ToDto().(*UserDTO)

	return &PostDTO{
		CommonDTO: common.CommonDTO{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		Title:   p.Title,
		Content: p.Content,
		Author:  *author,
	}
}

func (p PostDTO) ToEntity() common.Entity {
	author := p.Author.ToEntity().(*User)
	return &Post{
		CommonEntity: common.CommonEntity{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		Title:   p.Title,
		Content: p.Content,
		Author:  *author,
	}
}
