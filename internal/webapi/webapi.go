package webapi

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type TLSConfig struct {
	CertFile string
	KeyFile  string
	CAFile   string
}

func NewWebAPI() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func StartServer(e *echo.Echo, apiPort string, tlsConfig *TLSConfig) error {
	e.Logger.Infof("Starting HTTPS server on port: %s", apiPort)

	cert, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
	if err != nil {
		return fmt.Errorf("failed to load server certificate and key: %w", err)
	}

	var clientCAPool *x509.CertPool
	if tlsConfig.CAFile != "" {
		clientCAPool, err = buildClientCertPool(tlsConfig.CAFile)
		if err != nil {
			return fmt.Errorf("failed to build client CA pool: %w", err)
		}
	}

	e.Server.TLSConfig = buildTLSConfig(cert, clientCAPool)
	e.Server.Addr = fmt.Sprintf(":%s", apiPort)

	if err := e.Server.ListenAndServeTLS(tlsConfig.CertFile, tlsConfig.KeyFile); err != nil {
		return fmt.Errorf("failed to start HTTPS server: %w", err)
	}
	return nil
}

func buildClientCertPool(caFile string) (*x509.CertPool, error) {
	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to append CA cert")
	}

	return caCertPool, nil
}

func buildTLSConfig(cert tls.Certificate, caPool *x509.CertPool) *tls.Config {
	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	if caPool != nil {
		tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
		tlsCfg.ClientCAs = caPool
	}
	return tlsCfg
}
