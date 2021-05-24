package student

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/response"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/badoux/checkmail"
)

type StudentRequest struct {
	IdClass     int     `json:"id_class"`
	Sex         string  `json:"sex"`
	DateOfBirth string  `json:"date_of_birth"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	FullName    string  `json:"full_name"`
	CMND        string  `json:"cmnd"`
	CNSK        bool    `json:"cnsk"`
	GPLX        string  `json:"gplx"`
	Exp         int     `json:"exp"`
	NumberOfKm  int     `json:"number_of_km"`
	Amount      float64 `json:"amount"`
	Email       string  `json:"email"`
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
		resp, err := service.CreateStudent(req.Sex, req.DateOfBirth, req.Phone, req.Address,
			req.FullName, userRole.UserName, req.IdClass, req.CMND,
			req.CNSK, req.GPLX, req.Exp, req.NumberOfKm, req.Amount, req.Email)
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

	if s.CMND == "" {
		return errors.New("Bạn chưa nhập số chứng minh nhân dân")
	}
	if len(s.CMND) > 20 {
		return errors.New("Số chứng minh nhân dân không hợp lệ")
	}

	if len(s.GPLX) > 50 {
		return errors.New("Giấy phép lái xe không hợp lệ")
	}

	if s.Email != "" {
		if err := checkmail.ValidateFormat(s.Email); err != nil {
			return errors.New("Email không đúng định dạng")
		}
	}

	return nil
}
