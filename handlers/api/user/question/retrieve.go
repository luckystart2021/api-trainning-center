package question

import (
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func ShowSuiteTest(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "rank")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Hạng xe không tồn tại"))
			return
		}
		showTestSuite, err := service.GetAllTestSuiteByRank(code)
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
