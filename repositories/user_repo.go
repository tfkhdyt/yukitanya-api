package repositories

import "github.com/tfkhdyt/yukitanya-api/models"

type UserRepo interface {
	Store(user *models.User) error
	Show(userID uint) (*models.User, error)
	ShowByEmail(email string) (*models.User, error)
	ShowByUsername(username string) (*models.User, error)
}
