package request

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
)

type CreateTransaction struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int64   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func NewCreateTransaction(e echo.Context) (*CreateTransaction, error) {
	c := new(CreateTransaction)
	err := e.Bind(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *CreateTransaction) ToDomainModel() *model.Transaction {
	return &model.Transaction{
		AccountID: c.AccountID,
		Operation: model.OperationType(c.OperationTypeID),
		Amount:    c.Amount,
	}
}
