package common

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Err validator.ValidationErrorsTranslations
}

func NewValidationError(err validator.ValidationErrorsTranslations) *ValidationError {
	return &ValidationError{err}
}

func (v *ValidationError) Error() string {
	return "error"
}
