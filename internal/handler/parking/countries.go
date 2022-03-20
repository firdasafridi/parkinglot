package parking

import (
	"net/http"

	commonwriter "github.com/firdasafridi/parkinglot/lib/common/writer"
	"github.com/go-chi/chi"
)

// GetAllTransactionList get all transaction list brand
func (h *ParkingHandler) GetCounties(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	country := chi.URLParam(r, "country")
	listData, err := h.ParkingUC.GetCounties(ctx, country)
	if err != nil {
		commonwriter.WriteJSONAPIError(ctx, w, err)
		return
	}

	commonwriter.SetOKWithData(ctx, w, listData)
}
