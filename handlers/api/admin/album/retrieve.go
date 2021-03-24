package album

import (
	"api-trainning-center/service/admin/album"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetAlbums(service album.IAlbumService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showAlbums, err := service.GetListAlbum()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAlbums)
	}
}

func GetAlbum(service album.IAlbumService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã album không được rỗng"))
			return
		}

		idAlbum, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã album không hợp lệ"))
			return
		}
		data, err := service.GetAlbumDetail(idAlbum)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, data)
	}
}
