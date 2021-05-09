package fee

import (
	"api-trainning-center/service/admin/fee"
	"api-trainning-center/service/response"
	"net/http"
	"time"

	"github.com/leekchan/accounting"
)

type FeeResponse struct {
	ID        int       `json:"id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
}

func getFee(service fee.IFeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := FeeResponse{}
		showFee, err := service.ShowFees()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		ac := accounting.Accounting{Precision: 0}
		resp.ID = showFee.ID
		resp.CreatedAt = showFee.CreatedAt
		resp.UpdatedBy = showFee.UpdatedBy
		resp.Amount = ac.FormatMoney(showFee.Amount)
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
