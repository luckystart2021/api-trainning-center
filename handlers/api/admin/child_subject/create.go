package child_subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_subject"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

func createChildSubject(service child_subject.IChildSubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ChildSubject{}
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
		resp, err := service.CreateChildSubject(req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validateCreate(s models.ChildSubject) error {
	if len(s.Name) == 0 {
		return errors.New("Tên môn học chưa được nhập")
	}
	if len(s.Name) > 2000 {
		return errors.New("Tên môn học vượt quá ký tự giới hạn")
	}

	if s.Group.String == "0" {
		return errors.New("Nhóm môn học chưa được nhập")
	}

	if s.SubjectID == 0 {
		return errors.New("Mã môn học chưa được nhập")
	}
	return nil
}
