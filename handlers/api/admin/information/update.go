package information

import (
	"api-trainning-center/service/admin/information"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"api-trainning-center/validate"
	"errors"
	"net/http"
	"strconv"

	"github.com/badoux/checkmail"
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

		if err := validateUpdate(&req); err != nil {
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

func validateUpdate(c *InformationRequest) error {
	if c.Address == "" {
		return errors.New("Vui lòng nhập địa chỉ")
	}
	if len(c.Address) > 200 {
		return errors.New("Địa chỉ không được lớn hơn 200 ký tự")
	}

	if c.Phone == "" {
		return errors.New("Vui lòng nhập số điện thoại")
	}
	if len(c.Phone) > 15 || !validate.CheckNumber(c.Phone) {
		return errors.New("Số điện thoại không hợp lệ")
	}

	if c.Email == "" {
		return errors.New("Vui lòng nhập địa chỉ email")
	}
	if err := checkmail.ValidateFormat(c.Email); err != nil {
		return errors.New("Email không đúng định dạng")
	}

	if c.Maps == "" {
		return errors.New("Vui lòng nhập địa chỉ maps")
	}
	if len(c.Maps) > 2500 {
		return errors.New("Maps không được lớn hơn 1000 ký tự")
	}

	if c.Title == "" {
		return errors.New("Vui lòng nhập tiêu đề")
	}
	if len(c.Title) > 250 {
		return errors.New("Tiêu đề không được lớn hơn 250 ký tự")
	}

	if c.Description == "" {
		return errors.New("Vui lòng nhập miêu tả")
	}
	if len(c.Description) > 1000 {
		return errors.New("Miêu tả không được lớn hơn 1000 ký tự")
	}

	return nil
}
