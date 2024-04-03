package service

import (
	"context"
	"log/slog"
	"time"

	domainError "github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"

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
	slog.Debug("verifying if operation is valid", "operation_type_id", transaction.Operation)
	if !transaction.Operation.IsValid() {
		slog.Error("operation is not valid")
		return nil, domainError.OperationNotValid("operation is not valid")
	}

	slog.Debug("finding account present on transaction", "account", transaction.AccountID)
	_, err := c.accountFinder.Finder(ctx, transaction.AccountID)
	if err != nil {
		slog.Error("error while trying to find account present on transaction", "err", err)
		return nil, err
	}

	if !transaction.Operation.IsPayment() {
		transaction.Amount *= -1
	}

	transaction.EventDate = time.Now()

	slog.Debug("saving transaction on repository", "transaction", transaction)
	transaction, err = c.repository.Save(ctx, transaction)
	if err != nil {
		slog.Error("error to save transaction on repository", "err", err)
		return nil, err
	}

	slog.Info("transaction created successfully", "transaction", transaction)
	return transaction, nil

}
