package admin

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/account"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

// A completely separate router for administrator routes
func RouterLogin(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := account.NewStore(db)
	return func(router chi.Router) {
		router.Route("/admin", func(router chi.Router) {
			router.Post("/login", Login(st, client))
			router.With(middlewares.AuthJwtVerify).Post("/signup", CreateAccount(st))
			router.With(middlewares.AuthJwtVerify).Post("/change_password", CreateAccount(st))
		})
	}
}
