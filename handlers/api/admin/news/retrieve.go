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
		idChildArticle := chi.URLParam(r, "id_child_article")
		if idChildArticle == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không được rỗng"))
			return
		}
		idChildArticleP, err := strconv.Atoi(idChildArticle)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showArticles, err := service.ShowArticlesByChildCategory(idChildArticleP)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showArticles)
	}
}
