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
			router.Get("/view", RetrieveCourses(st))
			router.Get("/view/{id_course}", RetrieveCourse(st))
			router.Post("/create", CreateCourse(st))
		})
	}
}
