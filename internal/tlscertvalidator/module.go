package tlscertvalidator

import (
	"context"

	"github.com/ggsomnoev/mts-auth-service/internal/tlscertvalidator/handler"
	"github.com/labstack/echo/v4"
)

func Process(ctx context.Context, srv *echo.Echo) {
	handler.RegisterHandlers(ctx, srv)
}
