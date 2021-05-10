package teacher

import (
	"api-trainning-center/service/admin/teacher"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type TeacherR struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getTeachers(service teacher.ITeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showTeachers, err := service.ShowTeachers()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showTeachers)
	}
}

func getTeachersAvailable(service teacher.ITeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showTeachersAvalible, err := service.ShowTeacherByAvalible()
		if err != nil {
			response.RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		teachers := []TeacherR{}
		for _, data := range showTeachersAvalible {
			teacher := TeacherR{}
			teacher.Id = data.ID
			teacher.Name = data.Fullname
			teachers = append(teachers, teacher)
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, teachers)
	}
}

func getTeacher(service teacher.ITeacherService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã xe không được rỗng"))
			return
		}

		idTeacher, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã xe không hợp lệ"))
			return
		}
		showTeacher, err := service.ShowTeacher(idTeacher)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showTeacher)
	}
}
