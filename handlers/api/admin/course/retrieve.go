package course

import (
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func RetrieveCourses(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showCourses, err := service.ShowCoursesActive()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showCourses)
	}
}

func RetrieveCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id_course")
		idCourse, err := strconv.Atoi(id)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		showCourse, err := service.ShowCourses(idCourse)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showCourse)
	}
}
