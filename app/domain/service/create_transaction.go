package service

import (
	"context"
	"errors"
	"time"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
)

type TransactionCreator interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type CreatorTransaction struct {
	repository    repository.Transaction
	accountFinder AccountFinder
}

func NewTransactionCreator(transaction repository.Transaction, finder AccountFinder) *CreatorTransaction {
	return &CreatorTransaction{
		repository:    transaction,
		accountFinder: finder,
	}
}

func (c *CreatorTransaction) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	if !transaction.Operation.IsValid() {
		return nil, errors.New("transaction is not valid")
	}

	_, err := c.accountFinder.Finder(ctx, transaction.AccountID)
	if err != nil {
		return nil, err
	}

	if !transaction.Operation.IsPayment() {
		transaction.Amount *= -1
	}

	transaction.EventDate = time.Now()

	transaction, err = c.repository.Save(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil

}
