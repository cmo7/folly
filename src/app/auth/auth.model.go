package auth

import (
	"folly/src/lib/common"
)

// Authenticable is a representation of a authenticable entity in the database
// It contains the basic fields that all authenticable entities should have
// The Authenticable struct is embedded in other structs
// Usually, the Authenticable struct is embedded in the User struct
// Only fields that are common to all authenticable entities should be added here
// No relationships should be added here
type Authenticable struct {
	common.CommonEntity `gorm:"embedded"`
	Email               string `gorm:"unique;not null"`
	Password            string `gorm:"not null"`
}

type AuthenticableDTO struct {
	common.CommonDTO `json:",inline,omitempty"`
	Email            string `json:"email"`
}

type AuthenticableLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticableRegisterDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u Authenticable) ToDto() common.DTO {
	return &AuthenticableDTO{
		CommonDTO: common.CommonDTO{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email: u.Email,
	}
}

func (u AuthenticableDTO) ToEntity() common.Entity {
	return &Authenticable{
		CommonEntity: common.CommonEntity{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		Email: u.Email,
	}
}
