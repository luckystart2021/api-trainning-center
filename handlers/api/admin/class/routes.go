package class

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/class"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := class.NewStoreClass(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/class", func(router chi.Router) {
			router.Post("/create", CreateClass(st))
			router.Put("/{id}/update", UpdateClass(st))
			// router.Get("/{id_category}/views", RetrieveClass(st))
			// router.Get("/{id_category}/views", RetrieveChildCategories(st))
			// router.Get("/{id_child_category}/view/detail", RetrieveChildCategory(st))
			// router.Post("/create", CreateChildCategory(st))
			// router.Put("/{id}/update", UpdateChildCategory(st))
			// router.Put("/{id}/delete", DeleteChildCategory(st))
			// router.Put("/{id}/un-delete", UnDeleteChildCategory(st))
		})
	}
}
