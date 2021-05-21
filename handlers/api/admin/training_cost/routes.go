package training_cost

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/training_cost"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := training_cost.NewStoreCost(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/cost", func(router chi.Router) {
			router.Get("/{course_id}/views", GetCost(st))
			router.Get("/{class_id}/views", GetCostByClass(st))
			router.Post("/{course_id}/{class_id}/create", CreateCost(st))
			router.Put("/{cost_id}/{course_id}/{class_id}/update", UpdateCost(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", GetDetailCost(st))
				router.Delete("/delete", DeleteCost(st))
			})
		})
	}
}
