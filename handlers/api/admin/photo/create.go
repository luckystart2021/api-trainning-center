package photo

import (
	"api-trainning-center/middlewares"
	photoModel "api-trainning-center/models/admin/photo"
	"api-trainning-center/service/admin/photo"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"net/http"
)

func CreatePhoto(service photo.IPhotoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		userRole := r.Context().Value("values").(middlewares.Vars)
		form := r.MultipartForm
		files := form.File["upload[]"]
		for i, fileR := range files {
			file, err := files[i].Open()
			if err != nil {
				response.RespondWithError(w, http.StatusBadRequest, err)
				return
			}
			imgName, err := utils.FilesUpload(file, *fileR, "album")
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

			if imgName != "" {
				req.Img = imgName
			}

			_, err = service.CreatePhoto(req, userRole.UserName)
			if err != nil {
				response.RespondWithError(w, http.StatusBadRequest, err)
				return
			}

		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, response.MessageResponse{Status: true, Message: "Thêm ảnh thành công"})
	}
}
