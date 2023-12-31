package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/yukitanya-api/controllers/http"
	"github.com/tfkhdyt/yukitanya-api/middlewares"
)

func RegisterAuthRoute(app *fiber.App) {
	authController := di.GetInstance("authController").(*http.AuthController)
	auth := app.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Get("/inspect", middlewares.JwtMiddleware, authController.Inspect)
	auth.Post("/refresh", authController.RefreshToken)
}
