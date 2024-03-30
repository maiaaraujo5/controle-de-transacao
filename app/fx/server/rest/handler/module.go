package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/maiaaraujo5/controle-de-transacao/app/fx/service"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/handler"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(validator.New),
		service.Module(),
		createAccount(),
		findAccount(),
	)
}

func createAccount() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				handler.NewCreateAccount,
				fx.ResultTags(`group:"handlers"`),
				fx.As(new(handler.Handler))),
		),
	)
}

func findAccount() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				handler.NewFindAccount,
				fx.ResultTags(`group:"handlers"`),
				fx.As(new(handler.Handler))),
		),
	)
}
