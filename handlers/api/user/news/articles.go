package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetArticles(service article.IArticleService) http.HandlerFunc {
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
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticles)
	}
}

func GetArticle(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idArticle := chi.URLParam(r, "id_article")
		if idArticle == "" || len(idArticle) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã thông tin không tồn tại"))
			return
		}
		idArticleP, err := strconv.Atoi(idArticle)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã thông tin không hợp lệ"))
			return
		}
		meta := chi.URLParam(r, "meta")
		if meta == "" || len(meta) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã liên kết dữ liệu không tồn tại"))
			return
		}
		showArticle, err := service.ShowArticle(idArticleP, meta)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticle)
	}
}

func GetCategories(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_category_parent")
		if code == "" || len(code) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không tồn tại"))
			return
		}
		idCategoryParent, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showCategories, err := service.ShowCategories(idCategoryParent)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showCategories)
	}
}
