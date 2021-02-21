package slide

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/slide"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := slide.NewStoreSlide(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/slide", func(router chi.Router) {
			router.Get("/view", ShowSlides(st))
			router.Get("/{id}/view-detail", ShowDetailSlide(st))
			router.Post("/create", CreateSlide(st))
			router.Put("/{id}/update", UpdateSlide(st))
			// router.Put("/{id}/hide", HideSlide(st))
		})
	}
}
