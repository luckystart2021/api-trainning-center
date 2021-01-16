package account

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/user"

	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

// A completely separate router for administrator routes
func RouterLogin(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := user.NewStore(db)
	return func(router chi.Router) {
		router.Post("/system/login", Login(st, client))
		router.With(middlewares.AuthJwtVerify).Post("/change_password", ChangePassword(st, client))
		router.Group(adminRoute(db, client))
	}
}

func adminRoute(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := user.NewStore(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client))
		router.Route("/admin", func(router chi.Router) {
			router.Post("/signup", CreateAccount(st, client))
			router.Get("/logout", LogoutAccount(st, client))
			router.Post("/reset_password", ResetPassword(st, client))
			router.Get("/view/accounts", RetrieveAccounts(st, client))
			router.Get("/view/account/{username}", RetrieveAccount(st, client))
			router.Get("/delete/account/{username}", DeleteAccount(st, client))
		})
	}
}
