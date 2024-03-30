package dao

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/maiaaraujo5/controle-de-transacao/app/domain/model"
	dbModel "github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	"github.com/uptrace/bun"
)

type Account struct {
	db *bun.DB
}

func NewAccount(db *bun.DB) *Account {
	return &Account{
		db: db,
	}
}

func (a *Account) Save(ctx context.Context, account *model.Account) (*model.Account, error) {
	dbM := dbModel.NewAccountFromModel(account)

	_, err := a.db.NewInsert().Model(dbM).Exec(ctx)
	if err != nil {
		slog.Error("error to insert account on database", "err:", err)
		return nil, err
	}

	return dbM.ToDomainModel(), nil
}

func (a *Account) FindByID(ctx context.Context, ID int64) (*model.Account, error) {
	dbM := new(dbModel.Account)

	err := a.db.NewSelect().Model(dbM).Where("id=?", ID).Scan(ctx)
	if err != nil {
		slog.Error("error while trying to find account by id", "err:", err)
		return nil, err
	}

	return dbM.ToDomainModel(), nil
}

func (a *Account) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error) {
	dbM := new(dbModel.Account)

	err := a.db.NewSelect().Model(dbM).Where("document_number=?", documentNumber).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		slog.Error("error while trying to find account by document_number", "err:", err)
		return nil, err
	}

	return dbM.ToDomainModel(), nil
}
