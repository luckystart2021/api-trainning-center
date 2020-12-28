package api

import (
	"api-trainning-center/handlers/api/admin"

	"github.com/go-chi/chi"
)

func Route(router chi.Router) {
	//router.Group(admin.Router)
	router.Route("/admin", admin.Router)
}
