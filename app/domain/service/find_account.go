package service

import (
	"context"

	domainError "github.com/maiaaraujo5/controle-de-transacao/app/domain/errors"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
)

type AccountFinder interface {
	Finder(ctx context.Context, ID int64) (*model.Account, error)
}

type FinderAccount struct {
	repository repository.Account
}

func NewAccountFinder(account repository.Account) *FinderAccount {
	return &FinderAccount{
		repository: account,
	}
}

func (f *FinderAccount) Finder(ctx context.Context, ID int64) (*model.Account, error) {
	account, err := f.repository.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, domainError.AccountNotExists("no account was found")
	}

	return account, nil

}
