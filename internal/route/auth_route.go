package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	"github.com/tfkhdyt/yukitanya-api/internal/controller/http"
	"github.com/tfkhdyt/yukitanya-api/internal/middleware"
)

func RegisterAuthRoute(app *fiber.App) {
	authController := di.GetInstance("authController").(*http.AuthController)
	auth := app.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Get("/inspect", middleware.JwtMiddleware, authController.Inspect)
	auth.Post("/refresh", authController.RefreshToken)
}
