package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/teacher"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func updateTeacher(service teacher.ITeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã giáo viên không được rỗng"))
			return
		}

		idTeacher, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã giáo viên không hợp lệ"))
			return
		}

		req := models.Teacher{}
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
		resp, err := service.UpdateTeacher(idTeacher, req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
