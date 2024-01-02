package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = os.Getenv("JWT_KEY")

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(jwtKey)},
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	},
})
