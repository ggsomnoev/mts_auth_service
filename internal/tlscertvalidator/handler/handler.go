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
		if c.Request().TLS != nil && len(c.Request().TLS.PeerCertificates) > 0 {
			clientCert := c.Request().TLS.PeerCertificates[0]

			if clientCert.Subject.CommonName != "authorized-client" {
				return c.JSON(200, map[string]string{
					"message":     "Unauthorized: invalid certificate CN",
					"common_name": clientCert.Subject.CommonName,
				})
			}
		}

		return c.JSON(200, map[string]string{
			"message": "Authentication and authorization is successful",
		})
	}
}
