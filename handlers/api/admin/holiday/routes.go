package holiday

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/holiday"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := holiday.NewStoreHoliday(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/holiday", func(router chi.Router) {
			router.Get("/views", getHolidays(st))
			router.Post("/create", createHoliday(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", getHoliday(st))
				router.Put("/update", updateHoliday(st))
				router.Delete("/delete", deleteHoliday(st))
			})
		})
	}
}
