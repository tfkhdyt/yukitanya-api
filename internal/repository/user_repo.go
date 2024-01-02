package repository

import "github.com/tfkhdyt/yukitanya-api/internal/model"

type UserRepo interface {
	Store(user *model.User) error
	Show(userID uint) (*model.User, error)
	ShowByEmail(email string) (*model.User, error)
	ShowByUsername(username string) (*model.User, error)
}
