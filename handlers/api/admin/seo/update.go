package seo

import (
	modelSeo "api-trainning-center/models/admin/seo"
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateSeo(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := modelSeo.SeoRequest{}

		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã seo không được rỗng"))
			return
		}

		idSeo, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã seo không hợp lệ"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// if err := req.validate(); err != nil {
		// 	// If input is wrong, return an HTTP error
		// 	response.RespondWithError(w, http.StatusBadRequest, err)
		// 	return
		// }
		resp, err := service.UpdateSeo(idSeo, req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
