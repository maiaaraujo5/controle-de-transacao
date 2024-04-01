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

type CreateTransaction struct {
	service   service.TransactionCreator
	validator *validator.Validate
}

func NewCreateTransaction(service service.TransactionCreator, validate *validator.Validate) *CreateTransaction {
	return &CreateTransaction{
		service:   service,
		validator: validate,
	}
}

func (h *CreateTransaction) RegisterRoute(instance *echo.Echo) {
	instance.POST("/v1/transactions", h.Handle)
}

// Handle godoc
//
//	@Summary		Create an account
//	@Description	Create an account by given document number
//	@Accept			json
//	@Produce		json
//	@Param			account	body		request.CreateTransaction	true	"body for create one transaction"
//	@Success		201		{object}	response.Transaction
//	@Failure		400		{object}	middlewares.Error
//	@Failure		422		{object}	middlewares.ErrorValidation
//	@Failure		500		{object}	middlewares.Error
//	@Router			/v1/transaction [post]
func (h *CreateTransaction) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	req, err := request.NewCreateTransaction(c)
	if err != nil {
		slog.Error("error to bind create transaction request body", "err:", err)
		return err
	}

	err = h.validator.Struct(req)
	if err != nil {
		slog.Error("request body is not valid", "err:", err)
		return err
	}

	transaction, err := h.service.Create(ctx, req.ToDomainModel())
	if err != nil {
		slog.Error("error to create transaction", "err:", err)
		return err
	}

	slog.Debug("transaction created successfully")
	return c.JSON(http.StatusCreated, response.NewTransaction(transaction))
}
