package handlers

import (
	"api-trainning-center/handlers/api/admin"
	"api-trainning-center/utils"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// NewHandler create router
func NewHandler(db *sql.DB) http.Handler {
	router := chi.NewRouter()
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/api", admin.Router(db))
	return router
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
