package parking

import (
	"context"
	"net/http"

	parkinguc "github.com/firdasafridi/parkinglot/internal/usecase/parking"
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
