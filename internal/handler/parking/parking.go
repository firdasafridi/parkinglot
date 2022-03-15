package parking

import (
	"context"
	"net/http"

	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
	commonwriter "github.com/firdasafridi/parkinglot/lib/common/writer"
	"github.com/go-chi/chi"
)

type ParkingHandler struct {
	ParkingUC parkinguc.ParkingUC
}

// GetAllTransactionList get all transaction list brand
func (h *ParkingHandler) GetAllParkingTransactionList(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	listData, err := h.ParkingUC.GetAllParkingData(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, listData)
}

func (h *ParkingHandler) ParkVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	platNo := chi.URLParam(r, "platNo")
	err := h.ParkingUC.ParkVehicle(ctx, platNo)
	if err != nil {
		return
	}

	commonwriter.SetOKWithData(ctx, w, "ok")
}

func (h *ParkingHandler) LeaveParkingLot(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	platNo := chi.URLParam(r, "platNo")
	err := h.ParkingUC.LeaveParkingLot(ctx, platNo)
	if err != nil {
		return
	}

	commonwriter.SetOKWithData(ctx, w, "ok")
}
