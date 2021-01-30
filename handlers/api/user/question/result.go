package question

import (
	"api-trainning-center/models/admin/result"
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"
)

type Result struct {
	IdQuestion string `json:"id_question"`
	IdAnswer   string `json:"id_answer"`
}

func GetResult(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := []result.Result{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if len(req) == 0 {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Vui lòng chọn câu trả lời"))
			return
		}
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		showResult, err := service.ShowResult(req)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, showResult)
	}
}
