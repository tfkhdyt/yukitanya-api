package main

import (
	"log"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/goioc/di"
	_ "github.com/joho/godotenv/autoload"
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	routes.RegisterAuthRoute(app)

	if err := app.Listen(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatalln("Error:", err)
	}
}
