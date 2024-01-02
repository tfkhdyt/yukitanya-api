package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/common"
	"github.com/tfkhdyt/yukitanya-api/internal/dto"
	"github.com/tfkhdyt/yukitanya-api/internal/model"
	"github.com/tfkhdyt/yukitanya-api/internal/repository"
	"github.com/tfkhdyt/yukitanya-api/internal/repository/postgres"
)

type AuthUsecase struct {
	userRepo repository.UserRepo `di.inject:"userRepo"`
}

func NewAuthService(userRepo *postgres.UserRepoPg) *AuthUsecase {
	return &AuthUsecase{userRepo}
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

func (a *AuthUsecase) Login(payload *dto.LoginRequest) (*dto.LoginResponse, error) {
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
	accessToken, err := common.GenerateJWTToken(userID, common.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := common.GenerateJWTToken(userID, common.Refresh)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
