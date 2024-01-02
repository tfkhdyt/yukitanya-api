package service

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/matthewhartstonge/argon2"
)

type HashService struct {
	argon *argon2.Config `di.inject:"argon"`
}

func (h *HashService) HashPassword(password string) (string, error) {
	hashedPassword, err := h.argon.HashEncoded([]byte(password))
	if err != nil {
		log.Println("Error:", err)
		return "", fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	return string(hashedPassword), nil
}

func (h *HashService) VerifyPassword(password string, hashedPassword string) error {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashedPassword))
	if err != nil {
		log.Println("Error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to verify password")
	}

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid password")
	}

	return nil
}
