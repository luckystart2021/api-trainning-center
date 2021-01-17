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
		router.With(middlewares.AuthJwtVerify).Post("/account/change_password", ChangePassword(st, client))
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
			router.Route("/{username}", func(router chi.Router) {
				router.Get("/view/account", RetrieveAccount(st, client))
				router.Get("/disable/account", DisableAccount(st, client))
				router.Get("/enable/account", EnableAccount(st, client))
			})
			router.Post("/update/account", UpdateAccount(st, client))
		})
	}
}
