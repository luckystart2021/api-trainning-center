package child_subject

import (
	"api-trainning-center/service/admin/child_subject"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func getChildSubjects(service child_subject.IChildSubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "subject_id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã môn học không được rỗng"))
			return
		}

		idChildSubject, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã môn học không hợp lệ"))
			return
		}
		showSubjects, err := service.ShowChildSubjects(idChildSubject)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showSubjects)
	}
}

func getChildSubject(service child_subject.IChildSubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã môn học không được rỗng"))
			return
		}

		idSubject, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã môn học không hợp lệ"))
			return
		}
		showTeachers, err := service.ShowChildSubject(idSubject)
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showTeachers)
	}
}
