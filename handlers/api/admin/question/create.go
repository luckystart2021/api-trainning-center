package question

import (
	questionServeice "api-trainning-center/service/admin/question"
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"errors"
	"net/http"
	"strconv"
)

type QuestionRequest struct {
	CodeDe  string `json:"code_de"`
	Name    string `json:"name"`
	AnswerA string `json:"answer_a"`
	AnswerB string `json:"answer_b"`
	AnswerC string `json:"answer_c"`
	AnswerD string `json:"answer_d"`
	Img     string `json:"img"`
	Result  string `json:"result"`
	Liet    string `json:"liet"`
}

// CreateQuestion controller for creating new question
func CreateQuestion(service questionServeice.IQuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(32 << 20)
		imageName, err := utils.FileUpload(r, "question")
		//here we call the function we made to get the image and save it
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
			//checking whether any error occurred retrieving image
		}
		req := QuestionRequest{
			CodeDe:  r.FormValue("code_de"),
			Name:    r.FormValue("name"),
			AnswerA: r.FormValue("answer_a"),
			AnswerB: r.FormValue("answer_b"),
			AnswerC: r.FormValue("answer_c"),
			AnswerD: r.FormValue("answer_d"),
			Img:     imageName,
			Result:  r.FormValue("result"),
			Liet:    r.FormValue("liet"),
		}

		if err := validate(&req); err != nil {
			// If input is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		liet, err := strconv.ParseBool(req.Liet)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Không thể chuyển đổi string to boolean"))
			return
		}

		CodeDe, err := strconv.Atoi(req.CodeDe)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			response.RespondWithError(w, http.StatusBadRequest, errors.New("Không thể chuyển đổi string to int"))
			return
		}

		resp, err := service.CreateQuestion(CodeDe, req.Name, req.AnswerA, req.AnswerB, req.AnswerC, req.AnswerD, req.Img, req.Result, liet)
		if err != nil {
			response.RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		// send Result response
		response.RespondWithJSON(w, http.StatusOK, resp)
	}
}

func validate(q *QuestionRequest) error {
	if q.CodeDe == "" || len(q.CodeDe) == 0 {
		return errors.New("Vui lòng nhập mã đề")
	}
	if len(q.CodeDe) > 10 {
		return errors.New("Mã đề không hợp lệ")
	}

	if q.Name == "" || len(q.Name) == 0 {
		return errors.New("Vui lòng nhập câu hỏi")
	}
	if len(q.Name) > 2000 {
		return errors.New("Câu hỏi không hợp lệ")
	}

	if q.AnswerA == "" || len(q.AnswerA) == 0 {
		return errors.New("Vui lòng nhập câu trả lời A")
	}
	if len(q.AnswerA) > 2000 {
		return errors.New("Câu trả lời A không hợp lệ")
	}

	if q.AnswerB == "" || len(q.AnswerB) == 0 {
		return errors.New("Vui lòng nhập câu trả lời B")
	}
	if len(q.AnswerB) > 2000 {
		return errors.New("Câu trả lời B không hợp lệ")
	}

	if len(q.AnswerC) > 2000 {
		return errors.New("Câu trả lời C không hợp lệ")
	}
	if len(q.AnswerD) > 2000 {
		return errors.New("Câu trả lời D không hợp lệ")
	}

	if q.Result == "" || len(q.Result) == 0 {
		return errors.New("Vui lòng nhập đáp án")
	}
	if len(q.Result) > 2 {
		return errors.New("Đáp án không hợp lệ")
	}

	if len(q.Img) > 2000 {
		return errors.New("Hình ảnh không hợp lệ")
	}

	return nil
}
