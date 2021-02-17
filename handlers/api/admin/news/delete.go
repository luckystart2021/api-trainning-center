package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func DeleteArticle(service article.IArticleService) http.HandlerFunc {
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
		deleteArticleById, err := service.DeleteArticleById(idArticleP)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, deleteArticleById)
	}
}
