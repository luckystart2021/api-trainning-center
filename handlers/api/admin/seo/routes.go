package seo

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := seo.NewStoreSeo(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/seo", func(router chi.Router) {
			router.Put("/{id}/update", UpdateSeo(st))
			router.Get("/views", GetSeo(st))
		})
	}
}
