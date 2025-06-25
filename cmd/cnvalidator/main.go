package main

import (
	"github.com/ggsomnoev/mts-auth-service/internal/cnvalidator"
	"github.com/ggsomnoev/mts-auth-service/internal/config"
	"github.com/ggsomnoev/mts-auth-service/internal/webapi"
)

func main() {
	cfg := config.Load()

	srv := webapi.NewWebAPI()

	cnvalidator.Process(srv, cfg.TrustedClientCNs)

	tlsCfg := &webapi.TLSConfig{
		CertFile: cfg.WebAPICertFile,
		KeyFile:  cfg.WebAPIKeyFile,
		CAFile:   cfg.CACertFile,
	}

	srv.Logger.Fatal(webapi.StartServer(srv, cfg.Port, tlsCfg))
}
