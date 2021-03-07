package seo

import (
	"api-trainning-center/service/admin/seo"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		st := seo.NewStoreSeo(db)
		router.Get("/seo", GetSeo(st))
	}
}
