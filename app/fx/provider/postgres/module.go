package postgres

import (
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/repository"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/client"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/dao"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		accountModule(),
		transactionModule(),
		fx.Provide(
			client.NewOptions,
			client.NewClient,
		),
	)
}

func accountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				dao.NewAccount,
				fx.As(new(repository.Account)),
			),
		),
	)
}

func transactionModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				dao.NewTransaction,
				fx.As(new(repository.Transaction)),
			),
		),
	)
}
