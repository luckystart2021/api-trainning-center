package photo

import (
	"api-trainning-center/service/admin/photo"
	"api-trainning-center/service/response"

	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetPhotos(service photo.IPhotoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showPhotos, err := service.ShowPhotosInAdmin()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showPhotos)
	}
}

func GetPhoto(service photo.IPhotoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã ảnh không được rỗng"))
			return
		}

		idPhoto, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã ảnh không hợp lệ"))
			return
		}
		data, err := service.ShowPhotoInAdmin(idPhoto)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, data)
	}
}
