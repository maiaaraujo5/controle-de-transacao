package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/handler"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handlers []handler.Handler `group:"handlers"`
	Echo     *echo.Echo
}
