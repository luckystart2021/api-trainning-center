package photo

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/photo"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := photo.NewStorePhoto(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Route("/photo", func(router chi.Router) {
			router.Get("/views", GetPhotos(st))
			// router.Post("/create", CreateAlbum(st))
			router.Route("/{id}", func(router chi.Router) {
				router.Get("/view-detail", GetPhoto(st))
				// router.Put("/update", UpdateAlbum(st))
				// router.Delete("/delete", DeleteAlbum(st))
			})
		})
	}
}
