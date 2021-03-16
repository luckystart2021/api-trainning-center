package vehicle

import (
	"api-trainning-center/middlewares"
	models "api-trainning-center/models/admin/vehicle"
	"api-trainning-center/service/admin/vehicle"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

func UpdateVehicle(service vehicle.IVehicleService) http.HandlerFunc {
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

		req := models.VehicleUpdateRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := validateUpdate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateVehicle(idVehicle, req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateUpdate(req models.VehicleUpdateRequest) error {
	bsx := strings.TrimSpace(req.BienSoXe)
	if len(bsx) == 0 {
		return errors.New("Vui lòng nhập biển số xe")
	}
	if len(bsx) > 50 {
		return errors.New("Biển số xe không hợp lệ")
	}

	loaiXe := strings.TrimSpace(req.LoaiXe)
	if len(loaiXe) == 0 {
		return errors.New("Vui lòng nhập loại xe")
	}
	if len(loaiXe) > 100 {
		return errors.New("Loại xe không hợp lệ")
	}

	_, err := strconv.ParseBool(req.IsDeleted)
	if err != nil {
		return errors.New("Không đúng định dạng")
	}

	return nil
}
