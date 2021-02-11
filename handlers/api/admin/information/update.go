package information

import (
	"api-trainning-center/service/admin/information"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateInformation(service information.IInformationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		idInformation := chi.URLParam(r, "id_information")
		if idInformation == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã thông tin không được rỗng"))
			return
		}

		idInformationI, err := strconv.Atoi(idInformation)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã thông tin không hợp lệ"))
			return
		}

		imageName, err := utils.FileUpload(r, "information")
		//here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
			//checking whether any error occurred retrieving image
		}
		req := InformationRequest{
			Address:     r.FormValue("address"),
			Phone:       r.FormValue("phone"),
			Email:       r.FormValue("email"),
			Maps:        r.FormValue("maps"),
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}
		if imageName != "" || len(imageName) > 0 {
			req.Img = imageName
		}

		if err := validateInfo(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := service.UpdateInformation(idInformationI, req.Address, req.Phone, req.Email, req.Maps, req.Title, req.Description, req.Img)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
