package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/common"
	"github.com/tfkhdyt/yukitanya-api/dto"
	"github.com/tfkhdyt/yukitanya-api/models"
	"github.com/tfkhdyt/yukitanya-api/repositories"
	"github.com/tfkhdyt/yukitanya-api/repositories/postgres"
)

type UserService struct {
	userRepo repositories.UserRepo `di.inject:"userRepo"`
}

func NewUserService(userRepo *postgres.UserRepoPg) *UserService {
	return &UserService{userRepo}
}

func (u *UserService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if _, err := u.userRepo.ShowByEmail(payload.Email); err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email has been used")
	}

	if _, err := u.userRepo.ShowByUsername(payload.Username); err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Username has been used")
	}

	user := &models.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
	}

	var err error
	user.Password, err = common.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	if err := u.userRepo.Store(user); err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		Message: "User has been registered",
	}

	return response, nil
}

func (u *UserService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepo.ShowByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := common.VerifyPassword(payload.Password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := common.CreateJWTToken(user.ID, common.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := common.CreateJWTToken(user.ID, common.Refresh)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
