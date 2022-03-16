package parking

import (
	"context"
	"net/http"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
	"github.com/firdasafridi/parkinglot/lib/common"
	"github.com/firdasafridi/parkinglot/lib/common/commonerr"
	commonwriter "github.com/firdasafridi/parkinglot/lib/common/writer"
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

func (h *ParkingHandler) GetParkingHistoryByDate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
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
	ctx := context.Background()
	hstParking, err := h.ParkingUC.GetParkingHistoryDailyReport(ctx)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}
	commonwriter.SetOKWithData(ctx, w, hstParking)
}