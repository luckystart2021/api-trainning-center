package information

import (
	"api-trainning-center/service/admin/information"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := information.NewStoreInformation(db)
	return func(router chi.Router) {
		router.Get("/information", GetInformation(st))
		router.Post("/information/create", CreateInformation(st))
	}
}
