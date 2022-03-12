package main

import (
	"net/http"
	"time"

	"github.com/firdasafridi/parkinglot/internal/config"
	"github.com/firdasafridi/parkinglot/lib/common/log"
	"github.com/firdasafridi/parkinglot/lib/database"
)

func startServer(cfg *config.Config, handler http.Handler) error {
	srv := http.Server{
		ReadTimeout:  time.Duration(cfg.Server.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.HTTP.WriteTimeout) * time.Second,
		Handler:      handler,
	}

	log.Infoln("HTTP Serve", cfg.Server.HTTP.Address)
	return http.ListenAndServe(cfg.Server.HTTP.Address, srv.Handler)
}

func newDatabase(cfg *config.Config) *database.Connection {
	conn, err := database.New(appName, cfg.Database.DSN, database.Config{
		IsDebug: cfg.Database.Testing,
	})
	if err != nil {
		log.Fatalln("Failed to connect database", err)
	}
	return conn
}
