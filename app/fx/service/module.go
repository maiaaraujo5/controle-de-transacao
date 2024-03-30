package service

import (
	"github.com/maiaaraujo5/controle-de-transacao/app/domain/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/fx/provider/postgres"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		postgres.Module(),
		createAccountModule(),
		createTransactionModule(),
		findAccountModule(),
	)
}

func createAccountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewAccountCreator,
				fx.As(new(service.AccountCreator)),
			),
		),
	)
}

func createTransactionModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewTransactionCreator,
				fx.As(new(service.TransactionCreator)),
			),
		),
	)
}

func findAccountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewAccountFinder,
				fx.As(new(service.AccountFinder)),
			),
		),
	)
}
