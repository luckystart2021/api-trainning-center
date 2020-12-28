package handlers

import (
	"api-trainning-center/database"
	"api-trainning-center/handlers/api"
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
	router.Route("/api", api.Route)
	return router
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
