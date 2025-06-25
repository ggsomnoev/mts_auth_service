package main

import (
	"context"

	"github.com/ggsomnoev/mts-auth-service/internal/config"
	"github.com/ggsomnoev/mts-auth-service/internal/tlscertvalidator"
	"github.com/ggsomnoev/mts-auth-service/internal/webapi"
)

func main() {
	appCtx := context.Background()

	cfg := config.Load()

	srv := webapi.NewWebAPI()

	tlscertvalidator.Process(appCtx, srv)

	tlsCfg := &webapi.TLSConfig{
		CertFile: cfg.WebAPICertFile,
		KeyFile:  cfg.WebAPIKeyFile,
		CAFile:   cfg.CACertFile,
	}

	srv.Logger.Fatal(webapi.StartServer(srv, cfg.Port, tlsCfg))
}
