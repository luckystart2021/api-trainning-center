package question

import (
	"api-trainning-center/models/admin/result"
	"api-trainning-center/service/response"
	"database/sql"
)

type Question struct {
	Id           string    `json:"id"`
	QuestionName string    `json:"question_name"`
	Img          string    `json:"img"`
	Answers      []Answers `json:"answers"`
}

type Answers struct {
	Answer string `json:"answer"`
}

type ResponseResult struct {
	NumberOfCorrect   int          `json:"number_of_correct"`
	NumberOfIncorrect int          `json:"number_of_incorrect"`
	ResultTotal       string       `json:"result_total"`
	ResultTests       []ResultTest `json:"result_tests"`
}

type ResultTest struct {
	IdQuestion    string `json:"id_question"`
	IdAnswer      string `json:"id_answer"`
	CorrectAnswer string `json:"correct_answer"`
}

type QuestionSystem struct {
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

type ResponseQuestionExam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type IQuestionService interface {
	ShowQuestions(code string) ([]Question, error)
	ShowResult(result []result.Result) (ResponseResult, error)
	CreateQuestion(codeDe int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error)
	ShowQuestionsSystem(code string) ([]QuestionSystem, error)
	ShowQuestionSystem(idQuestion string) (QuestionSystem, error)
	UpdateQuestion(idQuestion, codeDe int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error)
	ShowQuestionsExam() ([]ResponseQuestionExam, error)
	DeleteQuestion(idQuestion string) (response.MessageResponse, error)
}

type StoreQuestion struct {
	db *sql.DB
}

func NewStoreQuestion(db *sql.DB) *StoreQuestion {
	return &StoreQuestion{
		db: db,
	}
}
