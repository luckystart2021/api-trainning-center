package child_category

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_category"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateChildCategory(service child_category.IChildCategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := ChildCategoryRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã câu hỏi không được rỗng"))
			return
		}

		idP, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã câu hỏi không hợp lệ"))
			return
		}

		if err := req.validate(); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateChildCategory(idP, userRole.UserName, req.Title, req.Meta, req.IdCategory)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
