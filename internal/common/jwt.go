package common

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_KEY")

type JwtType uint

const (
	Access  JwtType = 0
	Refresh JwtType = 1
)

func GenerateJWTToken(userID uint, jwtType JwtType) (string, error) {
	var exp int64
	switch jwtType {
	case Access:
		exp = time.Now().Add(5 * time.Minute).Unix()
	case Refresh:
		exp = time.Now().Add(24 * time.Hour * 7).Unix()
	default:
		return "", fiber.NewError(fiber.StatusBadRequest, "Invalid JWT type")
	}

	claims := jwt.MapClaims{
		"id":  userID,
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		log.Println("Error:", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to signed jwt token")
	}

	return t, nil
}

func ExtractUserIDFromClaims(c *fiber.Ctx) (uint, error) {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to validate token")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to validate claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid user id type")
	}

	return uint(userID), nil
}

func ExtractUserIDFromJWTPayload(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtKey), nil
	})
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Failed to parse jwt payload")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid jwt claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, fiber.NewError(fiber.StatusBadRequest, "Invalid user id type")
	}

	return uint(userID), nil
}
