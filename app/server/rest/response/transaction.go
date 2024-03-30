package response

import (
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"time"
)

type Transaction struct {
	ID              int64     `json:"id"`
	AccountID       int64     `json:"account_id"`
	OperationTypeID int64     `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func NewTransaction(transaction *model.Transaction) *Transaction {
	return &Transaction{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.Operation.Index(),
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}
}
