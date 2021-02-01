package question

type UpdateQuestionRequest struct {
	Id      int64  `json:"id"`
	CodeDe  int64  `json:"code_de"`
	Name    string `json:"name"`
	AnswerA string `json:"answer_a"`
	AnswerB string `json:"answer_b"`
	AnswerC string `json:"answer_c"`
	AnswerD string `json:"answer_d"`
	Img     string `json:"img"`
	Result  string `json:"result"`
	Liet    bool   `json:"liet"`
}
