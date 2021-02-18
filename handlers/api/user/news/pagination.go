package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetArticlePagination(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_category")
		if code == "" || len(code) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không tồn tại"))
			return
		}
		idCategory, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}

		showArticles, err := service.ShowArticles(idCategory)
		if len(showArticles) == 0 {
			respPage := []utils.PageNumberResponse{}
			respPage = append(respPage, utils.PageNumberResponse{PageNumber: 0})
			response.RespondWithJSON(w, http.StatusOK, respPage)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		calculatePageNumber := utils.CalculatePageNumber(len(showArticles))

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, calculatePageNumber)
	}
}

func GetChildArticlesPagination(service article.IArticleService) http.HandlerFunc {
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
		if len(showChildArticles) == 0 {
			respPage := []utils.PageNumberResponse{}
			respPage = append(respPage, utils.PageNumberResponse{PageNumber: 0})
			response.RespondWithJSON(w, http.StatusOK, respPage)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		calculatePageNumber := utils.CalculatePageNumber(len(showChildArticles))
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, calculatePageNumber)
	}
}
