package question

import (
	questionServeice "api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func DeleteQuestion(service questionServeice.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idQuestion := chi.URLParam(r, "id_question")
		if idQuestion == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã câu hỏi không được rỗng"))
			return
		}
		deleteQuestion, err := service.DeleteQuestion(idQuestion)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, deleteQuestion)
	}
}
