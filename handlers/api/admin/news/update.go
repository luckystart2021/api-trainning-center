package news

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateArticle(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		idArticle := chi.URLParam(r, "id_article")
		if idArticle == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã bài viết không được rỗng"))
			return
		}
		idArticleP, err := strconv.Atoi(idArticle)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã bài viết không hợp lệ"))
			return
		}
		imageName, err := utils.FileUpload(r, "news")

		// here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		req := ArticleRequest{
			IdChildCategory: r.FormValue("id_child_category"),
			Title:           r.FormValue("title"),
			Description:     r.FormValue("description"),
			Details:         r.FormValue("details"),
			Meta:            r.FormValue("meta"),
			KeyWordSEO:      r.FormValue("key_word_seo"),
		}
		if imageName != "" || len(imageName) > 0 {
			req.Image = imageName
		}

		if err := validate(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		IdChildCategoryP, err := strconv.Atoi(req.IdChildCategory)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		userRole := r.Context().Value("values").(middlewares.Vars)
		resp, err := service.UpdateArticle(idArticleP, IdChildCategoryP, userRole.UserName, req.Title, req.Description, req.Details, req.Meta, req.KeyWordSEO, req.Image)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
