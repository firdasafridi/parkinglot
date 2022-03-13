package main

import (
	"github.com/firdasafridi/parkinglot/internal/config"
	metricsmw "github.com/firdasafridi/parkinglot/internal/handler/http/middlewares/metrics"
	parkinghandler "github.com/firdasafridi/parkinglot/internal/handler/parking"
	parkingdb "github.com/firdasafridi/parkinglot/internal/repo/db/parking"
	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
)

func app(cfg *config.Config) moduleHandler {
	dbConn := newDatabase(cfg)

	metricMiddleware := metricsmw.New(&metricsmw.Middleware{
		NewRelic: metricsmw.NewRelic{
			AppName:    cfg.NewRelic.AppName,
			LicenseKey: cfg.NewRelic.Secret,
		},
	})

	parkingDB := parkingdb.New(&parkingdb.ParkingDB{
		Conn: dbConn,
	})

	parkingUC := parkinguc.New(&parkinguc.Parking{
		ParkingDB: parkingDB,
	})

	parkingHandler := parkinghandler.ParkingHandler{
		ParkingUC: parkingUC,
	}

	return moduleHandler{
		ParkingHandler:   parkingHandler,
		MetricMiddleware: metricMiddleware,
	}
}
