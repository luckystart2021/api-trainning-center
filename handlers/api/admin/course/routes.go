package course

import (
	"api-trainning-center/service/admin/course"
	"database/sql"

	"github.com/go-chi/chi"
)

func CourseRoute(db *sql.DB) func(chi.Router) {
	st := course.NewStoreCourse(db)
	return func(router chi.Router) {
		router.Get("/view", RetrieveCourses(st))
		router.Get("/view/{id_course}", RetrieveCourse(st))
	}
}
