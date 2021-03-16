package vehicle

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/vehicle"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := vehicle.NewStoreVehicle(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/vehicle", func(router chi.Router) {
			router.Get("/views", GetVehicles(st))
			router.Post("/create", CreateVehicle(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", GetVehicle(st))
				router.Put("/update", UpdateVehicle(st))
			})
		})
	}
}
