package question

import (
	"api-trainning-center/models/admin/result"
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
	NumberOfCorrect   int    `json:"number_of_correct"`
	NumberOfIncorrect int    `json:"number_of_incorrect"`
	ResultTotal       string `json:"result_total"`
	//ResultTests       []ResultTest `json:"result_tests"`
}

type ResultTest struct {
}

type IQuestionService interface {
	ShowQuestions(code string) ([]Question, error)
	ShowResult(result []result.Result) (ResponseResult, error)
}

type StoreQuestion struct {
	db *sql.DB
}

func NewStoreQuestion(db *sql.DB) *StoreQuestion {
	return &StoreQuestion{
		db: db,
	}
}
