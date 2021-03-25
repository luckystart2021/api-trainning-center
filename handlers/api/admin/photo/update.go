package photo

import (
	"api-trainning-center/middlewares"
	photoModel "api-trainning-center/models/admin/photo"
	"api-trainning-center/service/admin/photo"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateAlbum(service photo.IPhotoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã bài viết không được rỗng"))
			return
		}
		idPhoto, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã bài viết không hợp lệ"))
			return
		}
		r.ParseMultipartForm(32 << 20)
		userRole := r.Context().Value("values").(middlewares.Vars)
		imageName, err := utils.FileUpload(r, "album")
		// here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		req := photoModel.PhotoRequest{
			IdAlbum: r.FormValue("id_album"),
			Title:   r.FormValue("title"),
			Meta:    r.FormValue("meta"),
		}

		if imageName != "" {
			req.Img = imageName
		}

		resp, err := service.UpdatePhoto(idPhoto, req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
