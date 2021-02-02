package question

import (
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func GetQuestionAnswer(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "id")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đề không tồn tại"))
			return
		}
		showQuestions, err := service.ShowQuestions(code)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showQuestions)
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
