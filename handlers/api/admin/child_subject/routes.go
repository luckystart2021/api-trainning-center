package child_subject

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_subject"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := child_subject.NewStoreChildSubject(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/child-subject", func(router chi.Router) {
			router.Get("/{subject_id}/views", getChildSubjects(st))
			router.Post("/create", createChildSubject(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", getChildSubject(st))
				router.Put("/update", updateChildSubject(st))
				router.Delete("/delete", deleteChildSubject(st))
			})
		})
	}
}
