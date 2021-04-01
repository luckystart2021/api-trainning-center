package student

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateStudent(service student.IStudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := StudentRequest{}
		id := chi.URLParam(r, "id_student")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã học viên không được rỗng"))
			return
		}

		idStudent, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã học viên không hợp lệ"))
			return
		}
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
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateStudent(idStudent, req.Sex, req.DateOfBirth, req.Phone, req.Address, req.FullName, userRole.UserName, req.IdClass, req.CMND, req.CNSK, req.GPLX, req.Exp, req.NumberOfKm)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
