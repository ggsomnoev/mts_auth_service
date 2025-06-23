package handler

import (
	"context"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(ctx context.Context, srv *echo.Echo) {
	srv.GET("/auth", handleTLSAuthValidation(ctx))
}

func handleTLSAuthValidation(ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "TLS certificate validation endpoint",
		})
	}
}
