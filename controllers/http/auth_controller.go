package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/common"
	"github.com/tfkhdyt/yukitanya-api/dto"
	"github.com/tfkhdyt/yukitanya-api/services"
)

type AuthController struct {
	userService *services.UserService `di.inject:"userService"`
}

func NewUserController(userService *services.UserService) *AuthController {
	return &AuthController{userService}
}

func (a *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := common.ValidateBody(c, payload); err != nil {
		return err
	}

	response, err := a.userService.Register(payload)
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

	response, err := a.userService.Login(payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
