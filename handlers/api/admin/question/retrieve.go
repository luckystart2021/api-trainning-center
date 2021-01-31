package question

import (
	questionServeice "api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func ShowQuestions(service questionServeice.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code_de")
		if code == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã đề không tồn tại"))
			return
		}
		showQuestions, err := service.ShowQuestionsSystem(code)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showQuestions)
	}
}
