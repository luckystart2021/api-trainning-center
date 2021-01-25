package contact

import (
	"api-trainning-center/service/admin/contact"
	"api-trainning-center/service/response"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/badoux/checkmail"
)

type Contact struct {
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	Subject  string `json:"subject"`
}

func CreateContact(service contact.IContactService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req := Contact{}
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

		resp, err := service.CreateContact(req.FullName, req.Phone, req.Email, req.Message, req.Subject)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c Contact) validate() error {
	if c.FullName == "" {
		return errors.New("Vui lòng nhập họ và tên")
	}
	if len(c.FullName) > 150 {
		return errors.New("Họ và tên không hợp lệ")
	}

	if c.Phone == "" {
		return errors.New("Vui lòng nhập số điện thoại")
	}
	if len(c.Phone) > 15 || !validate.CheckNumber(c.Phone) {
		return errors.New("Số điện thoại không hợp lệ")
	}

	if len(c.Email) > 64 {
		return errors.New("Email không hợp lệ")
	}
	if c.Email != "" {
		if err := checkmail.ValidateFormat(c.Email); err != nil {
			return errors.New("Email không đúng định dạng")
		}
	}

	if c.Message == "" {
		return errors.New("Vui lòng nhập tin nhắn")
	}
	if len(c.Message) > 1000 {
		return errors.New("Tin nhắn không được lớn hơn 1000 ký tự")
	}

	if len(c.Subject) > 200 {
		return errors.New("Tiêu đề không được lớn hơn 200 ký tự")
	}
	return nil
}
