package student

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := student.NewStoreStudent(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/student", func(router chi.Router) {
			router.Get("/views", GetStudents(st))
			router.Post("/create", CreateStudent(st))
			router.Route("/{id_student}", func(router chi.Router) {
				router.Get("/view-detail", GetStudent(st))
				router.Put("/update", UpdateStudent(st))
				router.Delete("/delete", DeleteStudent(st))
			})
		})
	}
}
