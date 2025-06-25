package tlscertvalidator

import (
	"github.com/ggsomnoev/mts-auth-service/internal/tlscertvalidator/handler"
	"github.com/labstack/echo/v4"
)

func Process(srv *echo.Echo, trustedClientCNs []string) {
	handler.RegisterHandlers(srv, trustedClientCNs)
}
