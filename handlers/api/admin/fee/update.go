package fee

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/fee"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

func updateFee(service fee.IFeeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Fee{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := validateCreate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateFee(userRole.UserName, req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(req models.Fee) error {
	if req.Amount <= 0 {
		return errors.New("Số tiền chưa được nhập")
	}
	return nil
}
