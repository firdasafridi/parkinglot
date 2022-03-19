package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/firdasafridi/parkinglot/internal/handler/middleware"
	parkinghandler "github.com/firdasafridi/parkinglot/internal/handler/parking"
	"github.com/firdasafridi/parkinglot/lib/common/log"
	"github.com/firdasafridi/parkinglot/lib/common/writer"
)

type moduleHandler struct {
	ParkingHandler parkinghandler.ParkingHandler
	Middleware     *middleware.Config
}

func newRoutes(mHandler moduleHandler) *chi.Mux {

	log.Println("Starting to create new routing...")

	router := chi.NewRouter()

	// router.Use(mHandler.Middleware.HandlerNR)
	// router.Use(mHandler.Middleware.HandlerCustomMetricsTimeout)

	router.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer.WriteOK(r.Context(), w, "OK")
	}))

	router.Get("/parking/list", mHandler.ParkingHandler.GetAllParkingTransactionList)
	router.Get("/parking", mHandler.ParkingHandler.GetParkingLotByPlatNumber)
	router.Get("/parking/empty", mHandler.ParkingHandler.GetEmptyParkingLot)
	router.Get("/parking/history", mHandler.ParkingHandler.GetParkingHistoryByDate)
	router.Get("/parking/history/daily-report", mHandler.ParkingHandler.GetParkingHistoryDailyReport)

	return router
}
