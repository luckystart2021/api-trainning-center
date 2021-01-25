package information

import (
	"api-trainning-center/service/admin/information"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/badoux/checkmail"
)

type InformationRequest struct {
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Maps        string `json:"maps"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

func CreateInformation(service information.IInformationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		req := InformationRequest{}
		imageName, err := utils.FileUpload(r)
		//here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
			//checking whether any error occurred retrieving image
		}
		req.Img = imageName
		//fmt.Print("ssss", imageName)
		fmt.Println("Address", req)
		err = json.NewDecoder(r.Body).Decode(&req)
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

		resp, err := service.CreateInformation(req.Address, req.Phone, req.Email, req.Maps, req.Title, req.Description, req.Img)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c InformationRequest) validate() error {
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
	if len(c.Maps) > 1000 {
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

	if c.Img == "" {
		return errors.New("Vui lòng chọn hình ảnh")
	}
	if len(c.Img) > 250 {
		return errors.New("Đường dẫn hình ảnh không được lớn hơn 250 ký tự")
	}

	return nil
}
