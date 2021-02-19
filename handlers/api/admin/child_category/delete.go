package child_category

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/child_category"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteChildCategory(service child_category.IChildCategoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không được rỗng"))
			return
		}
		idR, err := strconv.Atoi(id)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		deleteCategoryById, err := service.DeleteCategoryById(idR, userRole.UserName)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, deleteCategoryById)
	}
}
