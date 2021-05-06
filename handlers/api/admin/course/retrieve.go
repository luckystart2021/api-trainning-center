package course

import (
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func RetrieveCourses(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_system")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không tồn tại"))
			return
		}
		showCourses, err := service.ShowCoursesActive(code)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showCourses)
	}
}

func RetrieveInActiveCourses(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showInActiveCourses, err := service.ShowCoursesInActive()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showInActiveCourses)
	}
}

func RetrieveCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_course")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không tồn tại"))
			return
		}
		showCourse, err := service.ShowCourses(code)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showCourse)
	}
}
