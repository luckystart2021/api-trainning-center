package child_subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_subject"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func updateChildSubject(service child_subject.IChildSubjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ChildSubject{}
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

		err = json.NewDecoder(r.Body).Decode(&req)
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
		resp, err := service.UpdateChildSubject(idSubject, req, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
