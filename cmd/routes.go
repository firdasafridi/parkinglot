package main

import (
	"net/http"

	"github.com/go-chi/chi"

	metricsmw "github.com/firdasafridi/parkinglot/internal/handler/http/middlewares/metrics"
	parkinghandler "github.com/firdasafridi/parkinglot/internal/handler/parking"
	"github.com/firdasafridi/parkinglot/lib/common/log"
	"github.com/firdasafridi/parkinglot/lib/common/writer"
)

type moduleHandler struct {
	ParkingHandler   parkinghandler.ParkingHandler
	MetricMiddleware *metricsmw.Middleware
}

func newRoutes(mHandler moduleHandler) *chi.Mux {

	log.Println("Starting to create new routing...")
	router := chi.NewRouter()
	router.Use(mHandler.MetricMiddleware.Handler)

	router.Get("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer.WriteOK(r.Context(), w, "OK")
	}))

	router.Get("/parking/list", mHandler.ParkingHandler.GetAllParkingTransactionList)

	return router
}
