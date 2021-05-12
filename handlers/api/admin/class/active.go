package class

import (
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
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không được rỗng"))
			return
		}

		idClass, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã khóa học không hợp lệ"))
			return
		}

		showAllClass, err := service.InActiveClass(idClass)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showAllClass)
	}
}
