package news

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"
)

type ArticleRequest struct {
	IdChildCategory string `json:"id_child_category"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Details         string `json:"details"`
	Meta            string `json:"meta"`
	Image           string `json:"image"`
	KeyWordSEO      string `json:"key_word_seo"`
}

func CreateArticle(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
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
		resp, err := service.CreateArticle(IdChildCategoryP, userRole.UserName, req.Title, req.Description, req.Details, req.Meta, req.KeyWordSEO, req.Image)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validate(a *ArticleRequest) error {
	if a.IdChildCategory == "" || len(a.IdChildCategory) == 0 {
		return errors.New("Vui lòng nhập mã danh mục")
	}
	if len(a.IdChildCategory) > 20 {
		return errors.New("Mã danh mục không hợp lệ")
	}

	if a.Title == "" || len(a.Title) == 0 {
		return errors.New("Vui lòng nhập tiêu đề bài viết")
	}
	if len(a.Title) > 4000 {
		return errors.New("Tiêu đề bài viết không được quá 4000 ký tự")
	}

	if a.Description == "" || len(a.Description) == 0 {
		return errors.New("Vui lòng nhập mô tả bài viết")
	}
	if len(a.Description) > 6000 {
		return errors.New("Mô tả bài viết không được quá 6000 ký tự")
	}

	if a.Details == "" || len(a.Details) == 0 {
		return errors.New("Vui lòng nhập chi tiết bài viết")
	}

	if a.Meta == "" || len(a.Meta) == 0 {
		return errors.New("Vui lòng nhập thẻ liên kết bài viết")
	}
	if len(a.Meta) > 4000 {
		return errors.New("Thẻ liên kết bài viết không được quá 4000 ký tự")
	}

	if a.KeyWordSEO == "" || len(a.KeyWordSEO) == 0 {
		return errors.New("Vui lòng nhập thẻ SEO bài viết")
	}
	if len(a.KeyWordSEO) > 4000 {
		return errors.New("Thẻ SEO không được quá 4000 ký tự")
	}

	return nil
}
