package question

import (
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"net/http"
)

func GetQuestionAnswer(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := "1"
		showQuestions, err := service.ShowQuestions(code)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showQuestions)
	}
}
