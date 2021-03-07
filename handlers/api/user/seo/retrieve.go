package seo

import (
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/response"
	"net/http"
)

func GetSeo(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showAbout, err := service.ShowSeos()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAbout)
	}
}
