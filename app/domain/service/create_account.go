package service

import (
	"context"
	"errors"
	"log/slog"

	domainError "github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
)

type AccountCreator interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)
}

type CreatorAccount struct {
	repository repository.Account
}

func NewAccountCreator(account repository.Account) *CreatorAccount {
	return &CreatorAccount{
		repository: account,
	}
}

func (a *CreatorAccount) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	slog.Debug("finding account with same document number on repository", "document_number", account.DocumentNumber)
	find, err := a.repository.FindByDocumentNumber(ctx, account.DocumentNumber)
	if err != nil {
		slog.Error("error while trying to find account with same document number", "err", err)
		return nil, errors.Join(err)
	}

	if find != nil {
		slog.Error("one account with this document number already exists")
		return nil, domainError.AccountAlreadyExists("an account with this document number already exists")
	}

	slog.Debug("saving account to database")
	account, err = a.repository.Save(ctx, account)
	if err != nil {
		slog.Error("error to save account", "error", err)
		return nil, err
	}

	slog.Info("account created successfully", "account", account)
	return account, nil
}
