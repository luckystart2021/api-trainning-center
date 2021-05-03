package subject

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/subject"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := subject.NewStoreSubject(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/subject", func(router chi.Router) {
			router.Get("/views", getSubjects(st))
			router.Post("/create", createSubject(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", getSubject(st))
				router.Put("/update", updateSubject(st, db))
			})
		})
	}
}
