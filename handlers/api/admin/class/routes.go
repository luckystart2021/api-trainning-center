package class

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/class"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := class.NewStoreClass(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/class", func(router chi.Router) {
			router.Post("/create", CreateClass(st))
			router.Put("/{id}/update", UpdateClass(st))
		})
	}
}
