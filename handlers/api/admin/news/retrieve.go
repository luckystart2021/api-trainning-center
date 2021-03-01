package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ShowArticles(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idChildArticle := chi.URLParam(r, "id_child_category")
		if idChildArticle == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không được rỗng"))
			return
		}
		idChildCategory, err := strconv.Atoi(idChildArticle)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showArticles, err := service.ShowArticlesByChildCategory(idChildCategory)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticles)
	}
}

func ShowArticlesDeleted(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showArticles, err := service.ShowArticlesDeleteByChildCategory()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticles)
	}
}

func ShowArticlesUnApproval(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showArticles, err := service.ShowArticlesUnApproval()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticles)
	}
}

func ShowDetailArticle(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idArticle := chi.URLParam(r, "id_article")
		if idArticle == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không được rỗng"))
			return
		}
		idArticleP, err := strconv.Atoi(idArticle)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showArticleById, err := service.ShowArticleById(idArticleP)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticleById)
	}
}
