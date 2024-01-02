package service

import (
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/yukitanya-api/internal/common"
)

type ValidatorService struct {
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidatorService() *ValidatorService {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalln("Error:", err)
	}

	return &ValidatorService{validate, trans}
}

func (v *ValidatorService) validateStruct(payload any) validator.ValidationErrorsTranslations {
	if err := v.validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)

		return errs.Translate(v.trans)
	}

	return nil
}

func (v *ValidatorService) ValidateBody(c *fiber.Ctx, payload any) error {
	if err := c.BodyParser(payload); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request body")
	}

	if err := v.validateStruct(payload); err != nil {
		return common.NewValidationError(err)
	}

	return nil
}
