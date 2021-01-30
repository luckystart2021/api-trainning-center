package about

import (
	"api-trainning-center/service/admin/about"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := about.NewStoreAbout(db)
	return func(router chi.Router) {
		router.Get("/about", GetAbout(st))
	}
}
