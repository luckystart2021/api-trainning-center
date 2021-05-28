package register

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/register"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := register.NewStoreRegister(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/register", func(router chi.Router) {
			router.Get("/{class_id}/views", GetRegisters(st))
			router.Get("/{id}/view-detail", GetRegisterDetail(st))
			router.Post("/create", Register(st))
			// router.Route("/{id_student}", func(router chi.Router) {
			// 	router.Get("/view-detail", GetStudent(st))
			// 	router.Put("/update", UpdateStudent(st))
			// 	router.Delete("/delete", DeleteStudent(st))
			// })
		})
	}
}
