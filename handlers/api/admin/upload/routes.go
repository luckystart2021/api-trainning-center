package upload

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/constant"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(client *redis.Client) func(chi.Router) {
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Post("/upload/ck", CkUpload())
	}
}
