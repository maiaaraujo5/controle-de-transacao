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

type CreateAccount struct {
	service   service.AccountCreator
	validator *validator.Validate
}

func NewCreateAccount(service service.AccountCreator, validate *validator.Validate) *CreateAccount {
	return &CreateAccount{
		service:   service,
		validator: validate,
	}
}

func (h *CreateAccount) RegisterRoute(instance *echo.Echo) {
	instance.POST("/v1/accounts", h.Handle)
}

// Handle godoc
//
//	@Summary		Create an account
//	@Description	Create an account by given document number
//	@Accept			json
//	@Produce		json
//	@Param			account	body		request.CreateAccount	true	"document number for account"
//	@Success		201		{object}	response.Account
//	@Failure		400		{object}	middlewares.Error
//	@Failure		422		{object}	middlewares.ErrorValidation
//	@Failure		500		{object}	middlewares.Error
//	@Router			/v1/accounts [post]
func (h *CreateAccount) Handle(e echo.Context) error {
	ctx := e.Request().Context()

	req, err := request.NewCreateAccount(e)
	if err != nil {
		slog.Error("error to bind create account request body", "err:", err)
		return err
	}

	err = h.validator.Struct(req)
	if err != nil {
		slog.Error("request body is not valid", "err:", err)
		return err
	}

	account, err := h.service.Create(ctx, req.ToDomainModel())
	if err != nil {
		slog.Error("error to create account", "err:", err)
		return err
	}

	slog.Debug("account created successfully")
	return e.JSON(http.StatusCreated, response.NewAccount(account))
}
