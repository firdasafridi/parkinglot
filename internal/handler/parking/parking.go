package parking

import (
	"context"
	"net/http"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
	"github.com/firdasafridi/parkinglot/lib/common"
	"github.com/firdasafridi/parkinglot/lib/common/commonerr"
	commonwriter "github.com/firdasafridi/parkinglot/lib/common/writer"
	"github.com/go-chi/chi"
)

type ParkingHandler struct {
	ParkingUC parkinguc.ParkingUC
}

// GetAllTransactionList get all transaction list brand
func (h *ParkingHandler) GetAllParkingTransactionList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

func (h *ParkingHandler) GetParkingLotByPlatNumber(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	platNo := r.URL.Query().Get("plat_no")

	data, err := h.ParkingUC.GetParkingLotByPlatNumber(ctx, platNo)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, data)
}

func (h *ParkingHandler) GetEmptyParkingLot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := h.ParkingUC.GetEmptyParkingLot(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, data)
}

func (h *ParkingHandler) GetParkingHistoryByDate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date := parkingdomain.ParkingDate{}
	err := common.DecodeSchema(r.URL.Query(), &date)
	if err != nil {
		err := commonerr.SetNewBadRequest("date", "invalid date params")
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}
	hstParking, err := h.ParkingUC.GetParkingHistoryByDate(ctx, date)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}
	commonwriter.SetOKWithData(ctx, w, hstParking)
}

func (h *ParkingHandler) GetParkingHistoryDailyReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hstParking, err := h.ParkingUC.GetParkingHistoryDailyReport(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, hstParking)
}
