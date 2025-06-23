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

	var tlsCfg *webapi.TLSConfig
	if cfg.Env != "local" {
		tlsCfg = &webapi.TLSConfig{
			CertFile: cfg.WebAPICertFile,
			KeyFile:  cfg.WebAPIKeyFile,
		}
	}

	srv.Logger.Fatal(webapi.StartServer(srv, cfg.Port, tlsCfg))
}
