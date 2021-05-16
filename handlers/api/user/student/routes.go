package student

import (
	"api-trainning-center/service/admin/student"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := student.NewStoreStudent(db)
	return func(router chi.Router) {
		router.Get("/student/search", getStudentByCode(st))
	}
}

func getStudentByCode(service student.IStudentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentCode := r.URL.Query().Get("code")
		if studentCode == "" || len(studentCode) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập mã học viên"))
			return
		}
		retrieveStudent, err := service.SearchStudentInformation(studentCode)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, retrieveStudent)
	}
}
