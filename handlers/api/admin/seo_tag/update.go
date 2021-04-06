package seo_tag

import (
	"api-trainning-center/service/admin/seo"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type SeoTagUpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UpdateSeoTag(service seo.ISeoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := []SeoTagUpdateRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp := response.MessageResponse{}
		for _, data := range req {
			t := strings.TrimSpace(data.Name)
			if len(t) == 0 {
				// If input is wrong, return an HTTP error
				response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập tên"))
				return
			}
			resp, err = service.UpdateSeoTags(data.Id, t)
			if err != nil {
				response.RespondWithError(w, http.StatusBadRequest, err)
				return
			}
			// send Result response
		}
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
