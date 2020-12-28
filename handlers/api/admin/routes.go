package admin

import "github.com/go-chi/chi"

func Router(r chi.Router) {
	r.Group(adminRoute)
}

func adminRoute(r chi.Router) {

}
