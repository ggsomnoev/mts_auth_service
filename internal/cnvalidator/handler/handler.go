package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(srv *echo.Echo, trustedClientCNs []string) {
	srv.GET("/auth", handleTLSAuthValidation(trustedClientCNs))
}

func handleTLSAuthValidation(trustedClientCNs []string) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().TLS == nil || len(c.Request().TLS.PeerCertificates) < 1 {
			return c.JSON(401, map[string]string{
				"message": "Unauthorized: no TLS client certificate provided",
			})
		}

		clientCN := c.Request().TLS.PeerCertificates[0].Subject.CommonName
		for _, trustedCN := range trustedClientCNs {
			if clientCN == trustedCN {
				return c.JSON(200, map[string]string{
					"message":     "Authentication and authorization successful",
					"common_name": clientCN,
				})
			}
		}

		return c.JSON(403, map[string]string{
			"message":     "Unauthorized: invalid certificate CN",
			"common_name": clientCN,
		})
	}
}
