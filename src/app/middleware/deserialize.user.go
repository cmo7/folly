package middleware

import (
	"folly/src/app/repositories"
	"folly/src/lib/common"
	"folly/src/lib/generics"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func DeserializeUser(c *fiber.Ctx) error {
	claims := c.Locals("claims").(jwt.MapClaims)

	id, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	user, err := repositories.UserRepositoryGORM.FindOne(id, common.NoConditions, []string{"Role"})
	if err != nil {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	c.Locals("user", user.ToDto())
	return c.Next()
}
