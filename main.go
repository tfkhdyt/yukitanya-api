package main

import (
	"errors"
	"log"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tfkhdyt/yukitanya-api/common"
	"github.com/tfkhdyt/yukitanya-api/controllers/http"
	"github.com/tfkhdyt/yukitanya-api/database"
	"github.com/tfkhdyt/yukitanya-api/repositories/postgres"
	"github.com/tfkhdyt/yukitanya-api/routes"
	"github.com/tfkhdyt/yukitanya-api/services"
)

func init() {
	_, _ = di.RegisterBean("userRepo", reflect.TypeOf((*postgres.UserRepoPg)(nil)))
	_, _ = di.RegisterBean("authService", reflect.TypeOf((*services.AuthService)(nil)))
	_, _ = di.RegisterBean("authController", reflect.TypeOf((*http.AuthController)(nil)))
	_, _ = di.RegisterBeanInstance("db", database.StartDB())
	_ = di.InitializeContainer()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			var ve *common.ValidationError
			if errors.As(err, &ve) {
				return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"error": ve.Err,
				})
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	routes.RegisterAuthRoute(app)

	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalln("Error:", err)
	}
}
