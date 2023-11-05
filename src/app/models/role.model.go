package models

import (
	"folly/src/database"
	"folly/src/lib/common"
)

func init() {
	database.RegisterMigration(&database.MigrationTask{
		Model:           &Role{},
		DropOnFlush:     true,
		TruncateOnFlush: true,
	})
}

type Role struct {
	common.CommonEntity `gorm:"embedded"`
	Name                string `gorm:"unique;not null"`
	Users               []User `gorm:"many2many:user_roles;"`
}

type RoleDTO struct {
	common.CommonDTO `json:",inline,omitempty"`
	Name             string    `json:"name"`
	Users            []UserDTO `json:"users,omitempty"`
}

func (r Role) ToDto() common.DTO {
	users := []UserDTO{}
	for _, user := range r.Users {
		users = append(users, *user.ToDto().(*UserDTO))
	}
	return &RoleDTO{
		CommonDTO: common.CommonDTO{
			ID:        r.ID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		},
		Name:  r.Name,
		Users: users,
	}
}

func (r RoleDTO) ToEntity() common.Entity {
	users := []User{}
	for _, user := range r.Users {
		users = append(users, *user.ToEntity().(*User))
	}
	return &Role{
		CommonEntity: common.CommonEntity{
			ID:        r.ID,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		},
		Name:  r.Name,
		Users: users,
	}
}
