package seo_tag

import (
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetSeoTag(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showAllSeoTag, err := service.ShowSeoTags()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllSeoTag)
	}
}

func GetDetailSeoTag(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không được rỗng"))
			return
		}

		idTag, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không hợp lệ"))
			return
		}
		showDetailSeoTag, err := service.ShowDetailSeoTags(idTag)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showDetailSeoTag)
	}
}
