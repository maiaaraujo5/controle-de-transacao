package echo

import (
	"context"

	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/middlewares"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			echoModule,
		),
		fx.Invoke(registerRoutes),
	)
}

func echoModule() *echo.Echo {
	e := echo.New()

	e.Use(middlewares.ErrorMiddleware)

	return e
}

func registerRoutes(params Params) {
	for _, handler := range params.Handlers {
		handler.RegisterRoute(params.Echo)
	}
}

func EchoLifeCycle(lc fx.Lifecycle, instance *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			g := new(errgroup.Group)

			g.Go(func() error {
				return instance.Start(":8080")
			})

			return g.Wait()
		},
		OnStop: func(ctx context.Context) error {
			return instance.Shutdown(ctx)
		},
	})
}
