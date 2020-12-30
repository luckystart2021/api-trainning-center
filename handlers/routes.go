package handlers

import (
	"api-trainning-center/database"
	"api-trainning-center/handlers/api/admin"
	"api-trainning-center/service/user"
	"api-trainning-center/utils"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var dbInstance database.Database

// NewHandler create router
func NewHandler(db database.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/api", adminRoute)
	return router
}

func adminRoute(router chi.Router) {
	router.Route("/admin", registerAdminRoute)
}

func registerAdminRoute(router chi.Router) {
	st := user.Store{dbInstance}
	router.Post("/signup", admin.CreateAccount(st))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, utils.ErrNotFound)
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, utils.ErrMethodNotAllowed)
}
