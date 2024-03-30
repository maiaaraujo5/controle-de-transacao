package service

import (
	"context"
	"errors"
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
	find, err := a.repository.FindByDocumentNumber(ctx, account.DocumentNumber)
	if err != nil {
		return nil, errors.Join(err)
	}

	if find != nil {
		return nil, domainError.AccountAlreadyExists("an account with this document number already exists")
	}

	account, err = a.repository.Save(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
