package seo_tag

import (
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type SeoTagRequest struct {
	Name string `json:"name"`
}

func CreateSeoTag(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := SeoTagRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		t := strings.TrimSpace(req.Name)
		if len(t) == 0 {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập tên"))
			return
		}
		resp, err := service.CreateSeoTag(req.Name)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
