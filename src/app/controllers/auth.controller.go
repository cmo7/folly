package controllers

import (
	"fmt"
	"folly/src/app/models"
	"folly/src/app/repositories"
	"folly/src/lib/generics"
	"folly/src/lib/helpers"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type LoginResponse struct {
	RefreshToken string         `json:"refresh_token"`
	AccessToken  string         `json:"access_token"`
	User         models.UserDTO `json:"user"`
}

type RegisterResponse struct {
	RefreshToken string         `json:"refresh_token"`
	AccessToken  string         `json:"access_token"`
	User         models.UserDTO `json:"user"`
}

var AuthController = struct {
	Login              func(*fiber.Ctx) error
	Register           func(*fiber.Ctx) error
	Logout             func(*fiber.Ctx) error
	RefreshAccessToken func(*fiber.Ctx) error
}{
	Login: func(c *fiber.Ctx) error {
		var payload LoginRequest

		if err := c.BodyParser(&payload); err != nil {
			return generics.BadRequest(c, err, "Invalid request")
		}

		errors := helpers.ValidateStruct(payload)
		if errors != nil {
			return generics.PayloadValidationFailed(c, errors, "Invalid request")
		}

		user, err := repositories.UserRepositoryGORM.FindByEmail(payload.Email, []string{"Roles"})
		if err != nil {
			return generics.Unauthorized(c, err, "Invalid credentials")
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
			return generics.Unauthorized(c, err, "Invalid credentials")
		}

		refreshToken, err := helpers.GenerateSignedToken(user.ID, helpers.RefreshToken)
		if err != nil {
			return generics.InternalServerError(c, err, "Failed to generate refresh token")
		}

		accessToken, err := helpers.GenerateSignedToken(user.ID, helpers.AccessToken)
		if err != nil {
			return generics.InternalServerError(c, err, "Failed to generate access token")
		}

		return generics.Ok(c, LoginResponse{
			RefreshToken: refreshToken,
			AccessToken:  accessToken,
			User:         *user.ToDto().(*models.UserDTO),
		}, "Login successful")

	},
	Register: func(c *fiber.Ctx) error {
		// Deserialize request body
		var payload RegisterRequest
		if err := c.BodyParser(&payload); err != nil {
			return generics.BadRequest(c, err, "Invalid request")
		}
		// Validate request body
		errors := helpers.ValidateStruct(payload)
		if errors != nil {
			return generics.PayloadValidationFailed(c, errors, "Invalid request")
		}
		// Check if email is taken
		if repositories.UserRepositoryGORM.IsEmailTaken(payload.Email) {
			return generics.BadRequest(c, fmt.Errorf("bad request"), "Email is already taken")
		}
		// Check if passwords match
		if payload.Password != payload.ConfirmPassword {
			return generics.BadRequest(c, nil, "Passwords do not match")
		}
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(payload.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return generics.InternalServerError(c, err, "Failed to hash password")
		}
		// Create user
		user := &models.User{
			Email:    payload.Email,
			Password: string(hashedPassword),
		}
		user, err = repositories.UserRepositoryGORM.Create(user)
		if err != nil {
			return generics.InternalServerError(c, err, "Failed to create user")
		}

		return generics.Created(c, RegisterResponse{
			User: *user.ToDto().(*models.UserDTO),
		}, "User created successfully")

	},
	Logout: func(c *fiber.Ctx) error {
		// Handled by frontend
		return nil

	},

	RefreshAccessToken: func(c *fiber.Ctx) error {
		// TODO: Implement refresh access token functionality
		return nil
	},
}
