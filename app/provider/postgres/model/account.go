package model

import (
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"table:accounts"`

	ID             int64  `bun:"id,pk,autoincrement"`
	DocumentNumber string `bun:"document_number"`
}

func NewAccountFromModel(account *model.Account) *Account {
	return &Account{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}

func (a *Account) ToDomainModel() *model.Account {
	return &model.Account{
		ID:             a.ID,
		DocumentNumber: a.DocumentNumber,
	}
}
