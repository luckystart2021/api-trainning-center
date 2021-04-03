package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/teacher"
	"api-trainning-center/service/response"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"
)

func CreateTeacher(service teacher.ITeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Teacher{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := validateCreate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.CreateTeacher(req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(s models.Teacher) error {
	if len(s.Sex) == 0 {
		return errors.New("Giới tính chưa được nhập")
	}
	if len(s.Sex) > 3 {
		return errors.New("Giới tính không hợp lệ")
	}

	if len(s.Dateofbirth) == 0 {
		return errors.New("Bạn chưa nhập ngày tháng năm sinh")
	}

	if len(s.Dateofbirth) > 10 || !validate.CheckDate(s.Dateofbirth) {
		return errors.New("Ngày tháng năm sinh không đúng định dạng")
	}

	if s.Phone == "" {
		return errors.New("Bạn chưa nhập số điện thoại")
	}
	if len(s.Phone) > 15 || !validate.CheckNumber(s.Phone) {
		return errors.New("Số điện thoại không đúng")
	}

	if s.Fullname == "" {
		return errors.New("Bạn chưa nhập họ và tên")
	}
	if len(s.Fullname) > 100 {
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

	if len(s.GPLX.String) > 50 {
		return errors.New("Giấy phép lái xe không hợp lệ")
	}
	return nil
}
