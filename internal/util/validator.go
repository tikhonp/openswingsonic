// Package util provide different helping functions and types for the project.
package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DefaultValidator struct {
	validator *validator.Validate
}

func (cv *DefaultValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		// TODO: fire error by opensabsonic api scheme
		return err
	}
	return nil
}

func NewDefaultValidator() echo.Validator {
	return &DefaultValidator{validator: validator.New()}
}
