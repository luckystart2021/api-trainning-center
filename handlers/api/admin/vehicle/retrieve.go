package vehicle

import (
	"api-trainning-center/service/admin/vehicle"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetVehicles(service vehicle.IVehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showVehicles, err := service.ShowVehicles()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showVehicles)
	}
}

func GetVehicle(service vehicle.IVehicleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã xe không được rỗng"))
			return
		}

		idVehicle, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã xe không hợp lệ"))
			return
		}
		showVehicle, err := service.ShowVehicle(idVehicle)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showVehicle)
	}
}
