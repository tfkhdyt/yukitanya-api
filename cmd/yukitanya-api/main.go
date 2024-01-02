package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tfkhdyt/yukitanya-api/internal/common"
	"github.com/tfkhdyt/yukitanya-api/internal/container"
	"github.com/tfkhdyt/yukitanya-api/internal/route"
)

func init() {
	container.InitializeContainer()
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
					"errors": ve.Errs,
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

	route.RegisterAuthRoute(app)

	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalln("Error:", err)
	}
}
