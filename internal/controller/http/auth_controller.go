package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/common"
	"github.com/tfkhdyt/yukitanya-api/internal/dto"
	"github.com/tfkhdyt/yukitanya-api/internal/service"
	"github.com/tfkhdyt/yukitanya-api/internal/usecase"
)

type AuthController struct {
	authUsecase  *usecase.AuthUsecase  `di.inject:"authUsecase"`
	tokenService *service.TokenService `di.inject:"tokenService"`
}

func NewUserController(userService *usecase.AuthUsecase, tokenService *service.TokenService) *AuthController {
	return &AuthController{userService, tokenService}
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := common.ValidateBody(c, payload); err != nil {
		fmt.Println(err.Error())
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
	if err := common.ValidateBody(c, payload); err != nil {
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
	if err := common.ValidateBody(c, payload); err != nil {
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
