package child_category

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_category"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := child_category.NewStoreChildCategory(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/child_category", func(router chi.Router) {
			router.Get("/{id_category}/view", RetrieveChildCategories(st))
			router.Post("/create", CreateChildCategory(st))
			router.Put("/{id}/update", UpdateChildCategory(st))
		})
	}
}
