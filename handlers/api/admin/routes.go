package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/account"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
)

// Router config
func Router(db *sql.DB) func(chi.Router) {
	return func(router chi.Router) {
		st := account.NewStore(db)
		router.Post("/login", Login(st))
	}
}

// A completely separate router for administrator routes
func CheckRouter(db *sql.DB) http.Handler {
	router := chi.NewRouter()
	router.Use(middlewares.AuthJwtVerify)
	st := account.NewStore(db)
	router.Post("/signup", CreateAccount(st))
	return router
}
