package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
)

type SearchResponse struct {
	Pages  []utils.PageNumberResponse `json:"pages"`
	Result []article.Article          `json:"result"`
}

func SearchNews(service article.IArticleService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := SearchResponse{}
		searchKey := r.URL.Query().Get("key")
		if searchKey == "" || len(searchKey) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng nhập từ khoá tìm kiếm"))
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
			resp.Result = showResultNewsByKey
			response.RespondWithJSON(w, http.StatusOK, resp)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		calculatePageNumber := utils.CalculatePageNumber(len(showResultNewsByKey))
		resp.Pages = calculatePageNumber
		resp.Result = showResultNewsByKey
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}
