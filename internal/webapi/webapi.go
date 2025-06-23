package webapi

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type TLSConfig struct {
	CertFile string
	KeyFile  string
}

func NewWebAPI() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func StartServer(e *echo.Echo, apiPort string, tlsConfig *TLSConfig) error {
	if tlsConfig != nil {
		log.Infof("Starting https server on port: %s", apiPort)
		if err := e.StartTLS(":"+apiPort, tlsConfig.CertFile, tlsConfig.KeyFile); err != nil {
			return fmt.Errorf("failed to start HTTPS server: %w", err)
		}
		return nil
	}

	log.Infof("Starting http server on port: %s", apiPort)
	if err := e.Start(":" + apiPort); err != nil {
		return fmt.Errorf("failed to start HTTP server: %w", err)
	}

	return nil
}
