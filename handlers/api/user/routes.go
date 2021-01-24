package user

import (
	"api-trainning-center/handlers/api/user/contact"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		router.Group(contact.Router(db))
	}
}
