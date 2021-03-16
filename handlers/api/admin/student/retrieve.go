package student

import (
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetStudents(service student.IStudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showStudents, err := service.ShowStudents()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showStudents)
	}
}

func GetStudent(service student.IStudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		showStudent, err := service.ShowStudent(idStudent)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showStudent)
	}
}
