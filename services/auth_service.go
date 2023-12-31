package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/common"
	"github.com/tfkhdyt/yukitanya-api/dto"
	"github.com/tfkhdyt/yukitanya-api/models"
	"github.com/tfkhdyt/yukitanya-api/repositories"
	"github.com/tfkhdyt/yukitanya-api/repositories/postgres"
)

type AuthService struct {
	userRepo repositories.UserRepo `di.inject:"userRepo"`
}

func NewAuthService(userRepo *postgres.UserRepoPg) *AuthService {
	return &AuthService{userRepo}
}

func (a *AuthService) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if _, err := a.userRepo.ShowByEmail(payload.Email); err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email has been used")
	}

	if _, err := a.userRepo.ShowByUsername(payload.Username); err == nil {
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

	if err := a.userRepo.Store(user); err != nil {
		return nil, err
	}

	response := &dto.RegisterResponse{
		Message: "User has been registered",
	}

	return response, nil
}

func (a *AuthService) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userRepo.ShowByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := common.VerifyPassword(payload.Password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := common.GenerateJWTToken(user.ID, common.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := common.GenerateJWTToken(user.ID, common.Refresh)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthService) Inspect(userID uint) (*dto.InspectResponse, error) {
	user, err := a.userRepo.Show(userID)
	if err != nil {
		return nil, err
	}

	return &dto.InspectResponse{
		ID:       user.Model.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
