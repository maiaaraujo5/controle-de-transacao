package service

import (
	"context"
	"log/slog"

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
	slog.Debug("finding account by id", "id", ID)
	account, err := f.repository.FindByID(ctx, ID)
	if err != nil {
		slog.Error("error to find account by id", "err", err)
		return nil, err
	}

	if account == nil {
		slog.Error("no account with this id was found", "id", ID)
		return nil, domainError.AccountNotExists("no account was found")
	}

	slog.Info("account successfully found", "account", account)
	return account, nil

}
