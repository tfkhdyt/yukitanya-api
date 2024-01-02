package common

import "strings"

type ValidationError struct {
	Errs []string
}

func NewValidationError(errs []string) *ValidationError {
	return &ValidationError{errs}
}

func (v *ValidationError) Error() string {
	return strings.Join(v.Errs, ",")
}
