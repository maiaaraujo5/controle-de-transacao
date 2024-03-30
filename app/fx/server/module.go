package server

import (
	"context"
	"log/slog"
	"os"

	"github.com/maiaaraujo5/controle-de-transacao/app/config"
	"github.com/maiaaraujo5/controle-de-transacao/app/fx/server/rest/echo"
	"github.com/maiaaraujo5/controle-de-transacao/app/fx/server/rest/handler"
	"go.uber.org/fx"
)

func Start() error {
	config.Load()

	app := fx.New(
		fx.Options(
			echo.Module(),
			handler.Module(),
			fx.Invoke(
				func() {
					slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
				},
			),
			fx.Invoke(echo.EchoLifeCycle),
		),
	)

	return app.Start(context.Background())
}
