package class

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/class"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

type ClassRequest struct {
	Name      string `json:"name"`
	IdCourse  int64  `json:"id_course"`
	Quantity  int64  `json:"quantity"`
	IdTeacher int64  `json:"id_teacher"`
}

func CreateClass(service class.IClassService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ClassRequest{}
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
		resp, err := service.CreateClass(userRole.UserName, req.Name, req.IdCourse, req.IdTeacher, req.Quantity)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func (c ClassRequest) validate() error {
	if c.Name == "" {
		return errors.New("Tên lớp chưa được nhập")
	}
	if len(c.Name) > 2000 {
		return errors.New("Tên lớp không hợp lệ")
	}

	if c.IdCourse == 0 {
		return errors.New("Mã khóa học chưa được nhập")
	}

	if c.Quantity == 0 {
		return errors.New("Số lượng học viên chưa được nhập")
	}

	if c.IdTeacher == 0 {
		return errors.New("Mã giáo viên chưa được nhập")
	}

	return nil
}
