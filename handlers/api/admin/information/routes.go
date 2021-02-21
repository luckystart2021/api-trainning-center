package information

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/information"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func InfoRouter(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := information.NewStoreInformation(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Get("/information/view", GetInformation(st))
		router.Post("/information/create", CreateInformation(st))
		router.Put("/information/{id_information}/update", UpdateInformation(st))
	}
}
