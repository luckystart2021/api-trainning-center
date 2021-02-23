package question

import (
	"api-trainning-center/models/admin/result"
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

func GetResult(service question.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := result.Result{}
		err := json.NewDecoder(r.Body).Decode(&req)
		logrus.WithFields(logrus.Fields{}).Info("[GetResult] request: ", req)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Định dạng dữ liệu không đúng, vui lòng thử lại"))
			return
		}

		if err := validate(req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
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

func validate(req result.Result) error {
	if len(req.Answers) == 0 {
		return errors.New("Vui lòng chọn câu trả lời")
	}
	if checkOverlap(req.Answers) {
		return errors.New("Vui lòng không chọn trùng câu hỏi")
	}
	if checkLenght(req.Answers) {
		return errors.New("Vui lòng chọn đáp án hợp lệ")
	}

	return nil
}

func checkLenght(as []result.Answer) bool {
	for _, data := range as {
		if len(data.IdAnswer) > 1 {
			return true
		}
	}
	return false
}

func checkOverlap(as []result.Answer) bool {
	for i := 0; i < len(as)-1; i++ {
		for j := i + 1; j < len(as); j++ {
			if as[i] == as[j] {
				return true
			}
		}
	}
	return false
}
