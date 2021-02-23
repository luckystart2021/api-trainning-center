package question

import (
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ShowRankVehicle(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showRank, err := service.GetAllRankVehicle()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showRank)
	}
}
func ShowSuiteTest(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_rank")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Hạng xe không tồn tại"))
			return
		}
		idRank, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã hạng xe không hợp lệ"))
			return
		}
		showTestSuite, err := service.GetAllTestSuiteByRank(idRank)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showTestSuite)
	}
}

func ShowQuestionsByIdSuite(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id_suite")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã bộ đề không tồn tại"))
			return
		}
		idSuite, err := strconv.Atoi(code)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã danh mục không hợp lệ"))
			return
		}
		showTestSuite, err := service.GetQuestionsByIdSuite(idSuite)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showTestSuite)
	}
}

func GetExam(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		showQuestionsExam, err := service.ShowQuestionsExam()
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showQuestionsExam)
	}
}
