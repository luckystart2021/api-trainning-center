package contact

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/contact"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := contact.NewStoreContact(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Get("/contact/view", ShowContacts(st))
	}
}
