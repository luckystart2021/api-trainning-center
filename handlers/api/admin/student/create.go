package student

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/response"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"
)

type StudentRequest struct {
	IdClass     int    `json:"id_class"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"date_of_birth"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	FullName    string `json:"full_name"`
}

func CreateStudent(service student.IStudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := StudentRequest{}
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
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.CreateStudent(req.Sex, req.DateOfBirth, req.Phone, req.Address, req.FullName, userRole.UserName, req.IdClass)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (s StudentRequest) validate() error {
	if len(s.Sex) == 0 {
		return errors.New("Giới tính chưa được nhập")
	}
	if len(s.Sex) > 3 {
		return errors.New("Giới tính không hợp lệ")
	}

	if len(s.DateOfBirth) == 0 {
		return errors.New("Bạn chưa nhập ngày tháng năm sinh")
	}

	if len(s.DateOfBirth) > 10 || !validate.CheckDate(s.DateOfBirth) {
		return errors.New("Ngày tháng năm sinh không đúng định dạng")
	}

	if s.Phone == "" {
		return errors.New("Bạn chưa nhập số điện thoại")
	}
	if len(s.Phone) > 15 || !validate.CheckNumber(s.Phone) {
		return errors.New("Số điện thoại không đúng")
	}

	if s.FullName == "" {
		return errors.New("Bạn chưa nhập họ và tên")
	}
	if len(s.FullName) > 100 {
		return errors.New("Họ và tên quá dài")
	}

	if s.Address == "" {
		return errors.New("Bạn chưa nhập địa chỉ")
	}
	if len(s.Address) > 500 {
		return errors.New("Địa chỉ quá dài")
	}

	return nil
}
