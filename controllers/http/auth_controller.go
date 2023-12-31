package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/common"
	"github.com/tfkhdyt/yukitanya-api/dto"
	"github.com/tfkhdyt/yukitanya-api/services"
)

type AuthController struct {
	authService *services.AuthService `di.inject:"authService"`
}

func NewUserController(userService *services.AuthService) *AuthController {
	return &AuthController{userService}
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := common.ValidateBody(c, payload); err != nil {
		return err
	}

	response, err := a.authService.Register(payload)
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

	response, err := a.authService.Login(payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (a *AuthController) Inspect(c *fiber.Ctx) error {
	userID, err := common.ExtractUserIDFromClaims(c)
	if err != nil {
		return err
	}

	response, err := a.authService.Inspect(uint(userID))
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

	userID, err := common.ExtractUserIDFromJWTPayload(payload.RefreshToken)
	if err != nil {
		return err
	}

	response, err := a.authService.RefreshToken(userID)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
