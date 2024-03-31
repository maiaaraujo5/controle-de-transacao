package middlewares

import (
	errors2 "errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"
)

type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ErrorValidation struct {
	Error
	ValidationError []ValidationError `json:"validationError,omitempty"`
}

type ValidationError struct {
	FieldPath string      `json:"path"`
	Field     string      `json:"field"`
	Value     interface{} `json:"value"`
	Message   string      `json:"message"`
}

func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			defer func() {
				err = switchError(c, err)
			}()
		}

		return err
	}
}

func switchError(c echo.Context, err error) error {
	switch {
	case errors.IsAccountAlreadyExists(err):
		return c.JSON(http.StatusConflict, newError(http.StatusConflict, err.Error()))
	case errors.IsOperationNotValid(err):
		return c.JSON(http.StatusPreconditionFailed, newError(http.StatusPreconditionFailed, err.Error()))
	case errors.IsAccountNotExists(err):
		return c.JSON(http.StatusNotFound, newError(http.StatusNotFound, err.Error()))
	case IsValidationError(err):
		return c.JSON(http.StatusUnprocessableEntity, newValidationError(err))
	default:
		return c.JSON(http.StatusInternalServerError, newError(http.StatusInternalServerError, err.Error()))
	}
}

func newError(code int, message string) *Error {
	return &Error{
		Code:        code,
		Description: message,
	}
}

func newValidationError(err error) *ErrorValidation {
	e := &ErrorValidation{
		Error: Error{
			Code:        http.StatusUnprocessableEntity,
			Description: "The server understands the content type of the request entity but was unable to process the contained instructions.",
		},
	}

	var validations []ValidationError
	for _, validation := range err.(validator.ValidationErrors) {
		validations = append(validations, ValidationError{
			FieldPath: validation.StructNamespace(),
			Field:     validation.Field(),
			Value:     validation.Value(),
			Message:   fmt.Sprintf("{%v} is a required field with type %v", validation.Field(), validation.Type().String()),
		})
	}

	e.ValidationError = validations
	return e
}

func IsValidationError(err error) bool {
	for err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return true
		}
		err = errors2.Unwrap(err)
	}

	return false
}
