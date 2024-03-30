package route

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/server/rest/handler"
)

func AddRoute(e *echo.Echo, handlers []handler.Handler) {
	for _, h := range handlers {
		h.RegisterRoute(e)
	}
}
