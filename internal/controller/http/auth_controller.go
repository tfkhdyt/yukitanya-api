package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/dto"
	"github.com/tfkhdyt/yukitanya-api/internal/service"
	"github.com/tfkhdyt/yukitanya-api/internal/usecase"
)

type AuthController struct {
	authUsecase      *usecase.AuthUsecase      `di.inject:"authUsecase"`
	tokenService     *service.TokenService     `di.inject:"tokenService"`
	validatorService *service.ValidatorService `di.inject:"validatorService"`
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := a.validatorService.ValidateBody(c, payload); err != nil {
		return err
	}

	response, err := a.authUsecase.Register(payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (a *AuthController) Login(c *fiber.Ctx) error {
	payload := new(dto.LoginRequest)
	if err := a.validatorService.ValidateBody(c, payload); err != nil {
		return err
	}

	response, err := a.authUsecase.Login(payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (a *AuthController) Inspect(c *fiber.Ctx) error {
	userID, err := a.tokenService.ExtractUserIDFromClaims(c)
	if err != nil {
		return err
	}

	response, err := a.authUsecase.Inspect(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (a *AuthController) RefreshToken(c *fiber.Ctx) error {
	payload := new(dto.RefreshTokenRequest)
	if err := a.validatorService.ValidateBody(c, payload); err != nil {
		return err
	}

	userID, err := a.tokenService.ExtractUserIDFromJWTPayload(payload.RefreshToken)
	if err != nil {
		return err
	}

	response, err := a.authUsecase.RefreshToken(userID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
