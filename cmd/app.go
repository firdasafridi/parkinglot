package main

import (
	"github.com/firdasafridi/parkinglot/internal/config"
	parkinghandler "github.com/firdasafridi/parkinglot/internal/handler/parking"
	parkingdb "github.com/firdasafridi/parkinglot/internal/repo/db/parking"
	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
	// TODO: #7.1 import package middleware here
)

func app(cfg *config.Config) moduleHandler {
	dbConn := newDatabase(cfg)
	parkingDB := parkingdb.New(&parkingdb.ParkingDB{
		Conn: dbConn,
	})

	parkingUC := parkinguc.New(&parkinguc.Parking{
		ParkingDB: parkingDB,
	})

	parkingHandler := parkinghandler.ParkingHandler{
		ParkingUC: parkingUC,
	}

	// TODO: #7 Init middleware here

	return moduleHandler{
		ParkingHandler: parkingHandler,
		// TODO: #8 Add middleware here
	}
}
