package about

import (
	"api-trainning-center/service/admin/about"
	"api-trainning-center/service/response"
	"net/http"
)

func GetAbout(service about.IAboutService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showAbout, err := service.ShowAbout()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAbout)
	}
}
