package training_cost

import (
	"api-trainning-center/service/admin/training_cost"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteCost(service training_cost.ICostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idCost := chi.URLParam(r, "id")
		if idCost == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không được rỗng"))
			return
		}
		costID, err := strconv.Atoi(idCost)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không hợp lệ"))
			return
		}
		resp, err := service.DeleteCost(costID)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
