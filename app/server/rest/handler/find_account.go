package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/request"
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

// Handle godoc
//
//	@Summary		Find an account
//	@Description	Find an account by given document number
//	@Accept			json
//	@Produce		json
//	@Param			accountId	path		int	true	"ID account"
//	@Success		200		{object}	response.Account
//	@Failure		400		{object}	middlewares.Error
//	@Failure		404		{object}	middlewares.Error
//	@Failure		422		{object}	middlewares.ErrorValidation
//	@Failure		500		{object}	middlewares.Error
//	@Router			/v1/accounts/{accountId} [get]
func (h *FindAccount) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	find, err := request.NewFindAccount(c)
	if err != nil {
		slog.Error("error to bind params", "err", err)
		return err
	}

	account, err := h.service.Finder(ctx, find.AccountID)
	if err != nil {
		slog.Error("error to find account", "err", err)
		return err
	}

	slog.Debug("find account successfully")
	return c.JSON(http.StatusOK, response.NewAccount(account))
}
