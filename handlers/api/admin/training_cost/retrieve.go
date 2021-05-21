package training_cost

import (
	"api-trainning-center/service/admin/training_cost"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetCost(service training_cost.ICostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		resp, err := service.ShowCost(courseID)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func GetCostByClass(service training_cost.ICostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "class_id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không được rỗng"))
			return
		}

		classID, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}
		resp, err := service.ShowCostByClass(classID)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func GetDetailCost(service training_cost.ICostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã chi phí không được rỗng"))
			return
		}

		costID, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã chi phí không hợp lệ"))
			return
		}
		resp, err := service.ShowDetailCost(costID)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}

}
