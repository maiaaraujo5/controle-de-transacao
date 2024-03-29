package repository

import (
	"context"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
)

type Transaction interface {
	Save(ctx context.Context, transaction *model.Transaction) error
	FindByID(ctx context.Context, ID string) (*model.Transaction, error)
}
