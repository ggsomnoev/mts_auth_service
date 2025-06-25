package cnvalidator

import (
	"github.com/ggsomnoev/mts-auth-service/internal/cnvalidator/handler"
	"github.com/labstack/echo/v4"
)

func Process(srv *echo.Echo, trustedClientCNs []string) {
	handler.RegisterHandlers(srv, trustedClientCNs)
}
