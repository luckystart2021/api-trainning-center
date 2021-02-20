package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/constant"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"
)

type SearchResponse struct {
	Pages  []utils.PageNumberResponse `json:"pages"`
	Result []ArticleResponse          `json:"result"`
}

func SearchNews(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := SearchResponse{}
		searchKey := r.URL.Query().Get("key")
		if searchKey == "" || len(searchKey) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập từ khoá tìm kiếm"))
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
		if len(searchKey) < 5 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập từ khoá tìm kiếm lớn hơn 5 ký tự"))
			return
		}
		showResultNewsByKey, err := service.ShowResultNewsByKey(searchKey)
		if len(showResultNewsByKey) == 0 {
			respPage := []utils.PageNumberResponse{}
			respPage = append(respPage, utils.PageNumberResponse{PageNumber: 0})
			resp.Pages = respPage
			resp.Result = []ArticleResponse{}
			response.RespondWithJSON(w, http.StatusOK, resp)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		respS, err := getDataSearchPage(pageNo, showResultNewsByKey)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		calculatePageNumber := utils.CalculatePageNumber(len(showResultNewsByKey))
		resp.Pages = calculatePageNumber
		resp.Result = respS
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

// pages start at 1, can't be 0 or less.
func getDataSearchPage(page int, data []article.Article) ([]ArticleResponse, error) {
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
