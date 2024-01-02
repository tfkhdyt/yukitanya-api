package container

import (
	"context"
	"log"
	"reflect"

	"github.com/goioc/di"
	"github.com/matthewhartstonge/argon2"
	"github.com/tfkhdyt/yukitanya-api/internal/controller/http"
	"github.com/tfkhdyt/yukitanya-api/internal/database"
	"github.com/tfkhdyt/yukitanya-api/internal/repository/postgres"
	"github.com/tfkhdyt/yukitanya-api/internal/service"
	"github.com/tfkhdyt/yukitanya-api/internal/usecase"
)

func InitializeContainer() {
	_, _ = di.RegisterBeanInstance("db", database.StartDB())
	if _, err := di.RegisterBeanFactory("argon", di.Singleton, func(ctx context.Context) (any, error) {
		argonInstance := argon2.DefaultConfig()

		return &argonInstance, nil
	}); err != nil {
		log.Fatalln("Error (argon):", err)
	}

	if _, err := di.RegisterBean("hashService", reflect.TypeOf((*service.HashService)(nil))); err != nil {
		log.Fatalln("Error (hashService):", err)
	}
	_, _ = di.RegisterBean("tokenService", reflect.TypeOf((*service.TokenService)(nil)))
	_, _ = di.RegisterBeanInstance("validatorService", service.NewValidatorService())

	_, _ = di.RegisterBean("userRepo", reflect.TypeOf((*postgres.UserRepoPg)(nil)))
	_, _ = di.RegisterBean("authUsecase", reflect.TypeOf((*usecase.AuthUsecase)(nil)))
	_, _ = di.RegisterBean("authController", reflect.TypeOf((*http.AuthController)(nil)))

	if err := di.InitializeContainer(); err != nil {
		log.Fatalln("Error:", err)
	}
}
