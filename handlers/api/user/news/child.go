package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetChildArticles(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metaParent := chi.URLParam(r, "meta_parent")
		if metaParent == "" || len(metaParent) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã liên kết dữ liệu không tồn tại"))
			return
		}

		metaChild := chi.URLParam(r, "meta_child")
		if metaChild == "" || len(metaChild) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã liên kết dữ liệu không tồn tại"))
			return
		}

		idChildCategory := chi.URLParam(r, "id_child_category")
		if idChildCategory == "" || len(idChildCategory) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không tồn tại"))
			return
		}

		idChildCategoryP, err := strconv.Atoi(idChildCategory)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}

		showChildArticles, err := service.ShowChildArticles(idChildCategoryP, metaChild, metaParent)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showChildArticles)
	}
}
