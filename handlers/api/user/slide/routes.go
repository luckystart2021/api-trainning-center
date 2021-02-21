package slide

import (
	"api-trainning-center/service/admin/slide"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := slide.NewStoreSlide(db)
	return func(router chi.Router) {
		router.Get("/slide", GetSlides(st))
	}
}
