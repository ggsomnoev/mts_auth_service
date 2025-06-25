package main

import (
	"github.com/ggsomnoev/mts-auth-service/internal/config"
	"github.com/ggsomnoev/mts-auth-service/internal/tlscertvalidator"
	"github.com/ggsomnoev/mts-auth-service/internal/webapi"
)

func main() {
	cfg := config.Load()

	srv := webapi.NewWebAPI()

	tlscertvalidator.Process(srv, cfg.TrustedClientCNs)

	tlsCfg := &webapi.TLSConfig{
		CertFile: cfg.WebAPICertFile,
		KeyFile:  cfg.WebAPIKeyFile,
		CAFile:   cfg.CACertFile,
	}

	srv.Logger.Fatal(webapi.StartServer(srv, cfg.Port, tlsCfg))
}
