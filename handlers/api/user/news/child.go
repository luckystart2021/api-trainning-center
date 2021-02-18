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

type ChildArticlesResponse struct {
	Id          int64  `json:"id"`
	IsOne       bool   `json:"is_one"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Meta        string `json:"meta"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
}

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

		showChildArticles, err := service.ShowChildArticles(idChildCategoryP, metaChild, metaParent)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := getDataPageChildArticles(pageNo, showChildArticles)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

// pages start at 1, can't be 0 or less.
func getDataPageChildArticles(page int, data []article.ChildCategoryNewsList) ([]ChildArticlesResponse, error) {
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

	childArticlesResponse := []ChildArticlesResponse{}
	var isValueOne bool
	for i, data := range data[start:stop] {
		if i == 0 {
			isValueOne = true
		} else {
			isValueOne = false
		}
		childArticleResponse := ChildArticlesResponse{
			Id:          data.Id,
			IsOne:       isValueOne,
			Title:       data.Title,
			Description: data.Description,
			Img:         data.Img,
			Meta:        data.Meta,
			CreatedAt:   data.CreatedAt,
			CreatedBy:   data.CreatedBy,
		}
		childArticlesResponse = append(childArticlesResponse, childArticleResponse)
	}

	return childArticlesResponse, nil
}
