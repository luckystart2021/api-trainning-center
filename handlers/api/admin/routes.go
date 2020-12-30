package admin

import (
	"api-trainning-center/service/user"
	"database/sql"

	"github.com/go-chi/chi"
)

// Router config
func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		router.Route("/admin", func(r chi.Router) {
			st := user.NewStore(db)
			router.Post("/signup", CreateAccount(st))
			router.Post("/login", Login(st))
		})
	}
}
