package subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/subject"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

func createSubject(service subject.ISubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.Subject{}
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
		resp, err := service.CreateSubject(req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(s models.Subject) error {
	if len(s.Name) == 0 {
		return errors.New("Tên môn học chưa được nhập")
	}
	if len(s.Name) > 2000 {
		return errors.New("Tên môn học vượt quá ký tự giới hạn")
	}

	if s.Group == 0 {
		return errors.New("Số ngày học của môn không được bằng 0")
	}
	// 1 là B2, 2 là C, 3 là DEF
	if s.RankID.Int == 0 {
		return errors.New("Hạng xe đào tạo không hợp lệ")
	}

	if s.Type.Int >= 3 {
		return errors.New("Hạng xe đào tạo không hợp lệ")
	}
	// Type = 1 là môn học lý thuyết
	if s.Type.Int == 1 {
		if s.Time.Int == 0 {
			return errors.New("Số giờ học của môn không được bằng 0")
		}

		if !s.Time.Valid {
			return errors.New("Số giờ học của môn phải được nhập")
		}

		if s.TeacherID.Int == 0 {
			return errors.New("Tên giáo viên chưa được nhập")
		}

		if !s.TeacherID.Valid {
			return errors.New("Tên giáo viên chưa được nhập")
		}
	}

	// Type = 2 là môn học thực hành
	if s.Type.Int == 2 {
		if s.HourStudent.String == "" || s.HourStudent.String == "0" {
			return errors.New("Giờ/Học viên chưa được nhập")
		}

		if s.KMStudent.String == "" || s.KMStudent.String == "0" {
			return errors.New("Km/Học viên chưa được nhập")
		}

		if s.HourDateVehicle.Int == 0 {
			return errors.New("Giờ/Ngày xe chưa được nhập")
		}

		if !s.KMDateVehicle.Valid {
			return errors.New("Km/Ngày xe chưa được nhập")
		}
	}

	return nil
}
