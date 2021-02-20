package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/constant"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ArticleResponse struct {
	Id          int64  `json:"id"`
	IsOne       bool   `json:"is_one"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

func GetNews(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showNews, err := service.ShowNews()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showNews)
	}
}

func GetFavoriteNews(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showFavoriteNews, err := service.ShowFavoriteNews()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showFavoriteNews)
	}
}

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
		pageParam := r.URL.Query().Get("page")
		if pageParam == "" || len(pageParam) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập số trang"))
			return
		}
		pageNo, err := strconv.Atoi(pageParam)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Số trang không hợp lệ"))
			return
		}

		showArticles, err := service.ShowArticles(idCategory)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := getDataPage(pageNo, showArticles)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
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
		meta := chi.URLParam(r, "meta_article")
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

// pages start at 1, can't be 0 or less.
func getDataPage(page int, data []article.Article) ([]ArticleResponse, error) {
	if page == 0 {
		return nil, errors.New("Số trang không được bằng 0")
	}
	start := (page - 1) * constant.ItemsPerPage
	stop := start + constant.ItemsPerPage

	if start > len(data) {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}

	if stop > len(data) {
		stop = len(data)
	}

	if len(data[start:stop]) == 0 {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}

	articlesResponse := []ArticleResponse{}
	var isValueOne bool
	for i, data := range data[start:stop] {
		if i == 0 {
			isValueOne = true
		} else {
			isValueOne = false
		}
		articleResponse := ArticleResponse{
			Id:          data.Id,
			IsOne:       isValueOne,
			Title:       data.Title,
			Description: data.Description,
			Img:         data.Img,
			Meta:        data.Meta,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		articlesResponse = append(articlesResponse, articleResponse)
	}

	return articlesResponse, nil
}
