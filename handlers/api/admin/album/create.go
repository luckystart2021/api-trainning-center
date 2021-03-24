package album

import (
	"api-trainning-center/service/admin/album"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

type AlbumRequest struct {
	Name string `json:"name"`
	Meta string `json:"meta"`
}

func CreateAlbum(service album.IAlbumService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := AlbumRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := req.validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		resp, err := service.CreateAlbum(req.Name, req.Meta)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c AlbumRequest) validate() error {
	if c.Name == "" {
		return errors.New("Tên album chưa được nhập")
	}
	if len(c.Name) > 2000 {
		return errors.New("Tên album không hợp lệ")
	}

	if c.Meta == "" {
		return errors.New("Liên kết album chưa được nhập")
	}
	if len(c.Meta) > 2000 {
		return errors.New("Liên kết album không hợp lệ")
	}

	return nil
}
