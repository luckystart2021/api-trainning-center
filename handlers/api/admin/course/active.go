package course

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ActiveCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_course")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không tồn tại"))
			return
		}
		idCourse, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}

		userRole := r.Context().Value("values").(middlewares.Vars)

		resp, err := service.ActiveCourse(idCourse, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func InActiveCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_course")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không tồn tại"))
			return
		}
		idCourse, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}

		userRole := r.Context().Value("values").(middlewares.Vars)

		resp, err := service.InActiveCourse(idCourse, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
