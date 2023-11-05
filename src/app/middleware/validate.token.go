package middleware

import (
	"fmt"
	"folly/src/lib/generics"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func ValidateToken(c *fiber.Ctx) error {
	var token string

	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer") {
		token = strings.Split(authorization, " ")[1]
	} else if c.Cookies("token") != "" {
		token = c.Cookies("token")
	}

	if token == "" {
		return generics.Unauthorized(c, fmt.Errorf("invalid authorization"), "No token provided")
	}

	tokenBytes, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("jwt.secret")), nil
	})

	if err != nil {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	claims, ok := tokenBytes.Claims.(jwt.MapClaims)
	if !ok || !tokenBytes.Valid {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	if claims["sub"] == nil {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	if claims["iss"] == "" {
		return generics.Unauthorized(c, err, "Invalid token")
	}

	c.Locals("claims", claims)

	return c.Next()
}
