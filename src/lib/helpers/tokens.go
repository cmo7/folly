package helpers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type TokenKind string

const (
	AccessToken  TokenKind = "access"
	RefreshToken TokenKind = "refresh"
)

// GenerateSignedToken generates a signed token with the user id as subject
func GenerateSignedToken(id uuid.UUID, kind TokenKind) (string, error) {

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": viper.GetString("jwt.issuer"),
		"sub": id,
		"exp": time.Now().Add(viper.GetDuration("jwt.exp")).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	})

	// Sign token with secret
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))

	return tokenString, err
}

func GenerateTokenCookie(token string) *fiber.Cookie {
	// Get token config

	// Create cookie
	cookie := &fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(viper.GetDuration("jwt.max_age").Seconds()),
		HTTPOnly: true,
		Domain:   "localhost",
	}

	return cookie
}
