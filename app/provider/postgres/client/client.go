package client

import (
	"database/sql"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewClient(options *Options) (*bun.DB, error) {

	sqlDb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(options.URL),
		pgdriver.WithUser(options.Username),
		pgdriver.WithPassword(options.Password),
		pgdriver.WithDatabase(options.Database),
		pgdriver.WithInsecure(true),
	))

	b := bun.NewDB(sqlDb, pgdialect.New())

	err := b.Ping()
	if err != nil {
		slog.Error("error to connect on database", "err:", err)
		return nil, err
	}

	return b, nil
}
