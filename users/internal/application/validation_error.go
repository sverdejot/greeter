package application

import (
	"errors"
	"fmt"
)

const (
	validationMsg = "some validation errors occured"
)

type ValidationError struct {
	error
}

func NewValidationError(errs ...error) ValidationError {
	return ValidationError{fmt.Errorf("%s: %w", validationMsg, errors.Join(errs...))}
}
