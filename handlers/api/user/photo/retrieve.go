package photo

import (
	"api-trainning-center/service/admin/photo"
	"api-trainning-center/service/response"

	"net/http"
)

func GetPhotos(service photo.IPhotoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showPhotos, err := service.ShowPhotos()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showPhotos)
	}
}
