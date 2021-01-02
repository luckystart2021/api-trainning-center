package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/account"
	"database/sql"

	"github.com/go-chi/chi"
)

// Router config
func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		st := account.NewStore(db)
		router.Use(middlewares.SetContentTypeMiddleware)
		router.Post("/login", Login(st))
		router.Use(middlewares.AuthJwtVerify)
		router.Post("/signup", CreateAccount(st))
	}
}
