package postgres

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/model"
	"gorm.io/gorm"
)

type UserRepoPg struct {
	db *gorm.DB `di.inject:"db"`
}

func (u *UserRepoPg) Store(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create new user")
	}

	return nil
}

func (u *UserRepoPg) ShowByEmail(email string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.First(user, "email = ?", email).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User is not found")
	}

	return user, nil
}

func (u *UserRepoPg) ShowByUsername(username string) (*model.User, error) {
	user := new(model.User)
	if err := u.db.First(user, "username = ?", username).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User is not found")
	}

	return user, nil
}

func (u *UserRepoPg) Show(userID uint) (*model.User, error) {
	user := new(model.User)
	if err := u.db.First(user, userID).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "User is not found")
	}

	return user, nil
}
