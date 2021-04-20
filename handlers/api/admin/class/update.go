package class

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/class"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ClassUpdateRequest struct {
	IdCourse  int64 `json:"id_course"`
	Quantity  int64 `json:"quantity"`
	IdDeleted bool  `json:"is_deleted"`
	TeacherId int64 `json:"teacher_id"`
	VehicleId int64 `json:"vehicle_id"`
}

func UpdateClass(service class.IClassService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ClassUpdateRequest{}

		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không được rỗng"))
			return
		}

		idClass, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không hợp lệ"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := req.validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateClass(idClass, userRole.UserName, req.IdCourse, req.Quantity, req.IdDeleted, req.TeacherId, req.VehicleId)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c ClassUpdateRequest) validate() error {
	if c.IdCourse == 0 {
		return errors.New("Mã khóa học chưa được nhập")
	}

	if c.Quantity == 0 {
		return errors.New("Số lượng học viên chưa được nhập")
	}

	return nil
}
