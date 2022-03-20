package main

import (
	"net/http"

	"github.com/go-chi/chi"

	parkinghandler "github.com/firdasafridi/parkinglot/internal/handler/parking"
	"github.com/firdasafridi/parkinglot/lib/common/log"
	"github.com/firdasafridi/parkinglot/lib/common/writer"
	// TODO: #3 import handler middleware here
)

type moduleHandler struct {
	ParkingHandler parkinghandler.ParkingHandler

	// TODO: #4 Setter handler middleware here
}

func newRoutes(mHandler moduleHandler) *chi.Mux {

	log.Println("Starting to create new routing...")
	router := chi.NewRouter()

	// TODO: #5 Set handler middleware here

	router.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer.WriteOK(r.Context(), w, "OK")
	}))

	router.Get("/parking/list", mHandler.ParkingHandler.GetAllParkingTransactionList)
	router.Post("/parking/park/{platNo}", mHandler.ParkingHandler.ParkVehicle)
	router.Post("/parking/leave/{platNo}", mHandler.ParkingHandler.LeaveParkingLot)
	router.Get("/parking", mHandler.ParkingHandler.GetParkingLotByPlatNumber)
	router.Get("/parking/empty", mHandler.ParkingHandler.GetEmptyParkingLot)
	router.Get("/parking/history", mHandler.ParkingHandler.GetParkingHistoryByDate)
	router.Get("/parking/history/daily-report", mHandler.ParkingHandler.GetParkingHistoryDailyReport)

	router.Get("/detail/{country}", mHandler.ParkingHandler.GetCounties)
	return router
}
