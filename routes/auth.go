package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/yukitanya-api/controllers/http"
)

func RegisterAuthRoute(app *fiber.App) {
	userController := di.GetInstance("userController").(*http.AuthController)
	auth := app.Group("/auth")

	auth.Post("/register", userController.Register)
	auth.Post("/login", userController.Login)
}
