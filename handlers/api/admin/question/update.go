package question

import (
	questionServeice "api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func UpdateQuestion(service questionServeice.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		idQuestion := chi.URLParam(r, "id_question")
		if idQuestion == "" {
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã câu hỏi không được rỗng"))
			return
		}

		idQuestionI, err := strconv.Atoi(idQuestion)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã câu hỏi không hợp lệ"))
			return
		}

		imageName, err := utils.FileUpload(r, "question")
		//here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
			//checking whether any error occurred retrieving image
		}
		req := QuestionRequest{
			Name:    r.FormValue("name"),
			Result:  r.FormValue("anwser_correct"),
			Liet:    r.FormValue("liet"),
			AnswerA: r.FormValue("answer_a"),
			AnswerB: r.FormValue("answer_b"),
			AnswerC: r.FormValue("answer_c"),
			AnswerD: r.FormValue("answer_d"),
		}
		if imageName != "" || len(imageName) > 0 {
			req.Img = imageName
		}
		if err := validate(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		liet, err := strconv.ParseBool(req.Liet)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Mã không hợp lệ"))
			return
		}

		updateQuestion, err := service.UpdateQuestion(idQuestionI, req.Name, req.AnswerA, req.AnswerB, req.AnswerC, req.AnswerD, req.Img, req.Result, liet)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, updateQuestion)
	}
}
