package dao

import (
	"context"
	"log/slog"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	dbModel "github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	"github.com/uptrace/bun"
)

type Transaction struct {
	db *bun.DB
}

func NewTransaction(db *bun.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}

func (t *Transaction) Save(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	dbM := dbModel.NewTransactionFromModel(transaction)

	_, err := t.db.NewInsert().Model(dbM).Exec(ctx)
	if err != nil {
		slog.Error("error to insert transaction on database", "err:", err)
		return nil, err
	}

	return dbM.ToDomainModel(), nil
}

func (t *Transaction) FindByID(ctx context.Context, ID string) (*model.Transaction, error) {
	//TODO implement me
	panic("implement me")
}
