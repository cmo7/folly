package models

import (
	"folly/src/database"
	"folly/src/lib/common"
)

func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &User{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}

type User struct {
	common.CommonEntity `gorm:"embedded"`
	Email               string `gorm:"unique;not null"`
	Password            string `gorm:"not null"`
	FirstName           string `gorm:"not null"`
	LastName            string `gorm:"not null"`
	Roles               []Role `gorm:"many2many:user_roles;"`
	Posts               []Post `gorm:"foreignKey:AuthorID"`
}

type UserDTO struct {
	common.CommonDTO `json:",inline,omitempty"`
	Email            string    `json:"email"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Roles            []RoleDTO `json:"roles,omitempty"`
	Posts            []PostDTO `json:"posts,omitempty"`
}

func (u User) ToDto() common.DTO {
	roles := []RoleDTO{}
	for _, role := range u.Roles {
		roles = append(roles, *role.ToDto().(*RoleDTO))
	}

	posts := []PostDTO{}
	for _, post := range u.Posts {
		posts = append(posts, *post.ToDto().(*PostDTO))
	}

	return &UserDTO{
		CommonDTO: common.CommonDTO{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Roles:     roles,
		Posts:     posts,
	}
}

func (u UserDTO) ToEntity() common.Entity {

	roles := []Role{}
	for _, role := range u.Roles {
		roles = append(roles, *role.ToEntity().(*Role))
	}

	posts := []Post{}
	for _, post := range u.Posts {
		posts = append(posts, *post.ToEntity().(*Post))
	}

	return &User{
		CommonEntity: common.CommonEntity{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Roles:     roles,
		Posts:     posts,
	}
}
