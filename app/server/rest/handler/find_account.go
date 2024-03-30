package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/response"
)

type FindAccount struct {
	service   service.AccountFinder
	validator *validator.Validate
}

func NewFindAccount(finder service.AccountFinder, validate *validator.Validate) *FindAccount {
	return &FindAccount{
		service:   finder,
		validator: validate,
	}
}

func (h *FindAccount) RegisterRoute(instance *echo.Echo) {
	instance.GET("/v1/accounts/:accountId", h.Handle)
}

func (h *FindAccount) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.ParseInt(c.Param("accountId"), 10, 64)
	if err != nil {
		return err
	}

	account, err := h.service.Finder(ctx, id)
	if err != nil {
		slog.Error("error to find account", "err:", err)
		return err
	}

	slog.Debug("find account successfully")
	return c.JSON(http.StatusOK, response.NewAccount(account))
}
