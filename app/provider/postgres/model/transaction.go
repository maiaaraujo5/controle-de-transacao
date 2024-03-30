package model

import (
	"time"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`

	ID        int64     `bun:"id,pk,autoincrement"`
	AccountID int64     `bun:"account_id"`
	Operation int64     `bun:"operation_type_id"`
	Amount    float64   `bun:"amount"`
	EventDate time.Time `bun:"event_date"`

	Account       *Account   `bun:"rel:belongs-to,join:account_id=id"`
	OperationType *Operation `bun:"rel:belongs-to,join:operation_type_id=id"`
}

func NewTransactionFromModel(transaction *model.Transaction) *Transaction {
	return &Transaction{
		ID:        transaction.ID,
		AccountID: transaction.AccountID,
		Operation: transaction.Operation.Index(),
		Amount:    transaction.Amount,
		EventDate: transaction.EventDate,
	}
}

func (t *Transaction) ToDomainModel() *model.Transaction {
	return &model.Transaction{
		ID:        t.ID,
		AccountID: t.AccountID,
		Operation: model.OperationType(t.Operation),
		Amount:    t.Amount,
		EventDate: t.EventDate,
	}
}
