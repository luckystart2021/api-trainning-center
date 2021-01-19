package course

import (
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"net/http"
)

func RetrieveCourse(service course.ICourseService) http.HandlerFunc {
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
