package photo

import (
	"api-trainning-center/service/admin/photo"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := photo.NewStorePhoto(db)
	return func(router chi.Router) {
		router.Get("/{id_album}/photos", GetPhotos(st))
	}
}
