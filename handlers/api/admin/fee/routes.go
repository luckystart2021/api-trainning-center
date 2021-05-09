package fee

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/fee"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := fee.NewStoreFee(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/fee", func(router chi.Router) {
			router.Get("/views", getFee(st))
			router.Put("/update", updateFee(st))
		})
	}
}
