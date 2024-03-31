package handler

import (
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

type Swagger struct {
}

func NewSwagger() *Swagger {
	return &Swagger{}
}

func (s *Swagger) RegisterRoute(instance *echo.Echo) {
	instance.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (s *Swagger) Handle(c echo.Context) error {
	return nil
}
