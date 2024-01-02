package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/dto"
	"github.com/tfkhdyt/yukitanya-api/internal/model"
	"github.com/tfkhdyt/yukitanya-api/internal/repository"
	"github.com/tfkhdyt/yukitanya-api/internal/service"
)

type AuthUsecase struct {
	userRepo     repository.UserRepo   `di.inject:"userRepo"`
	hashService  *service.HashService  `di.inject:"hashService"`
	tokenService *service.TokenService `di.inject:"tokenService"`
}

func (a *AuthUsecase) Register(payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if _, err := a.userRepo.ShowByEmail(payload.Email); err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email has been used")
	}

	if _, err := a.userRepo.ShowByUsername(payload.Username); err == nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Username has been used")
	}

	user := &model.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
	}

	var err error
	user.Password, err = a.hashService.HashPassword(payload.Password)
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

func (a *AuthUsecase) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := a.userRepo.ShowByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := a.hashService.VerifyPassword(payload.Password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := a.tokenService.GenerateJWTToken(user.ID, service.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.tokenService.GenerateJWTToken(user.ID, service.Refresh)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthUsecase) Inspect(userID uint) (*dto.InspectResponse, error) {
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

func (a *AuthUsecase) RefreshToken(userID uint) (*dto.RefreshTokenResponse, error) {
	accessToken, err := a.tokenService.GenerateJWTToken(userID, service.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.tokenService.GenerateJWTToken(userID, service.Refresh)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
