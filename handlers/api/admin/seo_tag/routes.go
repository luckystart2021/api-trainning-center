package seo_tag

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
		router.Route("/seo-tag", func(router chi.Router) {
			router.Get("/views", GetSeoTag(st))
			router.Post("/create", CreateSeoTag(st))
			router.Put("/update", UpdateSeoTag(st))
			router.Delete("/{id}/delete", DeleteSeoTag(st))
			router.Get("/{id}/view-detail", GetDetailSeoTag(st))
		})
	}
}
