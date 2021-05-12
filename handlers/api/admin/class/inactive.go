package class

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/class"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func InActiveClass(service class.IClassService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id_class")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không được rỗng"))
			return
		}

		idClass, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không hợp lệ"))
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		showAllClass, err := service.InActiveClass(idClass, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllClass)
	}
}

func ActiveClass(service class.IClassService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id_class")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không được rỗng"))
			return
		}

		idClass, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã lớp học không hợp lệ"))
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		showAllClass, err := service.ActiveClass(idClass, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllClass)
	}
}
