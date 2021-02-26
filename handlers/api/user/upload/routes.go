package upload

import "github.com/go-chi/chi"

func Router() func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/upload", Upload())
		router.Post("/upload/ck", CkUpload())
	}
}
