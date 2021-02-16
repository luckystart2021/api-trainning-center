package account

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/user"
	"api-trainning-center/service/constant"

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
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Post("/	", CreateAccount(st))
		router.Get("/logout", LogoutAccount(st, client))
		router.Post("/reset_password", ResetPassword(st))
		router.Get("/view/accounts", RetrieveAccounts(st))
		router.Route("/{username}", func(router chi.Router) {
			router.Get("/view/account", RetrieveAccount(st))
			router.Put("/disable/account", DisableAccount(st))
			router.Put("/enable/account", EnableAccount(st))
		})
		router.Put("/update/account", UpdateAccount(st))
	}
}
