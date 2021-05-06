package course

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/course"
	"api-trainning-center/service/response"
	"api-trainning-center/validate"
	"encoding/json"
	"errors"
	"net/http"
)

type Course struct {
	Name           string `json:"name"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	GraduationDate string `json:"graduation_date"`
	TestDate       string `json:"test_date"`
	TrainingSystem string `json:"training_system"`
	Time           string `json:"time"`
}

func CreateCourse(service course.ICourseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Course{}
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
		resp, err := service.CreateCourse(userRole.UserName, req.Name, req.StartDate, req.EndDate, req.GraduationDate, req.TestDate, req.TrainingSystem, req.Time)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c Course) validate() error {
	if c.Name == "" {
		return errors.New("Tên khóa học chưa được nhập")
	}
	if len(c.Name) > 250 {
		return errors.New("Mã khóa học không hợp lệ")
	}

	if c.TrainingSystem == "" {
		return errors.New("Hệ đào tạo chưa được nhập")
	}
	if len(c.TrainingSystem) > 50 {
		return errors.New("Hệ đào tạo không hợp lệ")
	}
	if !checkSystem(c.TrainingSystem) {
		return errors.New("Hệ đào tạo không hợp lệ")
	}

	if c.StartDate == "" {
		return errors.New("Ngày bắt đầu khóa học chưa được nhập")
	}
	if len(c.StartDate) > 10 || !validate.CheckDate(c.StartDate) {
		return errors.New("Ngày bắt đầu khóa học không đúng định dạng")
	}

	if c.EndDate == "" {
		return errors.New("Ngày kết thúc khóa học chưa được nhập")
	}
	if len(c.EndDate) > 10 || !validate.CheckDate(c.EndDate) {
		return errors.New("Ngày kết thúc khóa học không đúng định dạng")
	}

	if c.TestDate == "" {
		return errors.New("Ngày thi sát hạch chưa được nhập")
	}
	if len(c.TestDate) > 10 || !validate.CheckDate(c.TestDate) {
		return errors.New("Ngày thi sát hạch không đúng định dạng")
	}

	if len(c.GraduationDate) > 10 || !validate.CheckDate(c.GraduationDate) {
		return errors.New("Ngày tốt nghiệp khóa học không đúng định dạng")
	}

	if c.Time == "" {
		return errors.New("Thời gian học chưa được nhập")
	}
	if len(c.Name) > 20 {
		return errors.New("Thời gian học không hợp lệ")
	}

	return nil
}

func checkSystem(name string) bool {
	systemLst := map[string]string{
		"B2": "B2",
		"C":  "C",
		"D":  "D",
		"E":  "E",
		"F":  "F",
	}
	_, ok := systemLst[name]
	return ok
}
