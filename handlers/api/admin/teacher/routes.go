package teacher

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/teacher"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := teacher.NewStoreTeacher(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/teacher", func(router chi.Router) {
			router.Get("/views", getTeachers(st))
			router.Get("/views-available", getTeachersAvailable(st))
			router.Post("/create", CreateTeacher(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", getTeacher(st))
				router.Put("/update", updateTeacher(st))
			})
		})
	}
}
