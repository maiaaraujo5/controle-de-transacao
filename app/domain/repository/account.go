package repository

import (
	"context"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
)

type Account interface {
	Save(ctx context.Context, account *model.Account) (*model.Account, error)
	FindByID(ctx context.Context, ID int64) (*model.Account, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error)
}
