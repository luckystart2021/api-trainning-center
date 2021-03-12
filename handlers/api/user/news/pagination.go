package news

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetArticlePagination(service article.IArticleService) http.HandlerFunc {
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

		countArticles, err := service.CountArticles(idCategory)
		if countArticles == 0 {
			respPage := []utils.PageNumberResponse{}
			respPage = append(respPage, utils.PageNumberResponse{PageNumber: 0})
			response.RespondWithJSON(w, http.StatusOK, respPage)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		calculatePageNumber := utils.CalculatePageNumber(countArticles)

		// send Result response
		response.RespondWithJSON(w, http.StatusOK, calculatePageNumber)
	}
}

func GetChildArticlesPagination(service article.IArticleService) http.HandlerFunc {
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

		countChildArticles, err := service.CountChildArticles(metaChild, metaParent)
		if countChildArticles == 0 {
			respPage := []utils.PageNumberResponse{}
			respPage = append(respPage, utils.PageNumberResponse{PageNumber: 0})
			response.RespondWithJSON(w, http.StatusOK, respPage)
			return
		}
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		calculatePageNumber := utils.CalculatePageNumber(countChildArticles)
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, calculatePageNumber)
	}
}
