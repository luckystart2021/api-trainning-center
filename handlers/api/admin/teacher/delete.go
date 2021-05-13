package teacher

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/teacher"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func inActive(service teacher.ITeacherService) http.HandlerFunc {
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

		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.InActive(idTeacher, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func active(service teacher.ITeacherService) http.HandlerFunc {
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

		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.Active(idTeacher, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
