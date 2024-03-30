package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"
	"net/http"
)

type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
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
