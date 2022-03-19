package main

import (
	"os"

	"github.com/firdasafridi/parkinglot/internal/config"
	"github.com/firdasafridi/parkinglot/lib/common/log"

	// TODO: #6.1 import pacakge nr here
	"github.com/firdasafridi/parkinglot/lib/util/nr"
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

	// TODO: #6 Call new package nr here
	err = nr.New(cfg.NewRelic.AppName, cfg.NewRelic.Secret, map[string]string{
		"env": os.Getenv("ENV"),
	})
	if err != nil {
		log.Errorln(err)
	}

	mHandler := app(cfg)

	httpServer := newRoutes(mHandler)

	log.Errorln(startServer(cfg, httpServer))
}
