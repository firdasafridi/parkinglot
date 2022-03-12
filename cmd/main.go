package main

import (
	"github.com/firdasafridi/parkinglot/internal/config"
	"github.com/firdasafridi/parkinglot/lib/common/log"
)

const (
	appName = "parking-lot"
)

func main() {
	log.Infoln("Starting new service...")

	cfg, err := config.New("parkinglot")
	if err != nil {
		log.Fatalln("Can't get config file", err)
	}

	mHandler := app(cfg)

	httpServer := newRoutes(mHandler)

	log.Errorln(startServer(cfg, httpServer))
}
