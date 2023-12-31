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

func (u *AuthController) Register(c *fiber.Ctx) error {
	payload := new(dto.RegisterRequest)
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	if err := common.Validate(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	response, err := u.userService.Register(payload)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
