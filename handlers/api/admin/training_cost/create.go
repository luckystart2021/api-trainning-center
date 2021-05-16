package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/training_cost"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CreateCost(service training_cost.ICostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.TrainingCost{}
		id := chi.URLParam(r, "course_id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không được rỗng"))
			return
		}

		courseID, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}

		idC := chi.URLParam(r, "class_id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không được rỗng"))
			return
		}

		classID, err := strconv.Atoi(idC)
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
		if err := validateCreate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.CreateCost(req, courseID, classID, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(s models.TrainingCost) error {
	if s.Amount == 0 {
		return errors.New("Vui lòng nhập số tiền")
	}
	return nil
}
