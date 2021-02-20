package course

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func CourseRoute(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := course.NewStoreCourse(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/course", func(router chi.Router) {
			router.Get("/view/active", RetrieveCourses(st))
			router.Get("/view/in-active", RetrieveInActiveCourses(st))
			router.Post("/create", CreateCourse(st))
			router.Get("/{id_course}/view", RetrieveCourse(st))
			router.Put("/{id_course}/update", UpdateCourse(st))
		})
	}
}
