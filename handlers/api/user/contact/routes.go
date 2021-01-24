package contact

import (
	"api-trainning-center/service/admin/contact"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := contact.NewStoreContact(db)
	return func(router chi.Router) {
		router.Post("/contact/create", CreateContact(st))
	}
}
