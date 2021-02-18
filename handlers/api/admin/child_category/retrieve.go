package child_category

import (
	"api-trainning-center/service/admin/child_category"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func RetrieveChildCategories(service child_category.IChildCategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id_category")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không được rỗng"))
			return
		}

		idCategory, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showChildCategories, err := service.ShowChildCategories(idCategory)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showChildCategories)
	}
}
