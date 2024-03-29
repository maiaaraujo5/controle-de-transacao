package request

import (
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
)

type CreateTransaction struct {
	AccountID       int64   `json:"account_id"`
	OperationTypeID int64   `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func (c *CreateTransaction) ToDomainModel() *model.Transaction {
	return &model.Transaction{
		AccountID:   c.AccountID,
		OperationID: model.OperationType(c.OperationTypeID),
		Amount:      c.Amount,
	}
}
