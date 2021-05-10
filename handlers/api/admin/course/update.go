package course

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateCourse(service course.ICourseService) http.HandlerFunc {
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
		req := Course{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := req.validateUpdate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateCourse(idCourse, userRole.UserName, req.Name, req.StartDate, req.EndDate, req.GraduationDate, req.TestDate, req.TrainingSystem, req.Time)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
