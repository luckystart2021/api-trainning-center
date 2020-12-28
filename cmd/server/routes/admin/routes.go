package admin

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Users(router chi.Router) {
	router.Get("/login", retieveUser)
}

func retieveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Hello world`)
}
