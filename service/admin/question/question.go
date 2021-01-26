package question

import "database/sql"

type Question struct {
	Id           string    `json:"id"`
	QuestionName string    `json:"question_name"`
	Img          string    `json:"img"`
	Answers      []Answers `json:"answers"`
}

type Answers struct {
	Answer string `json:"answer"`
}

type IQuestionService interface {
	ShowQuestions(code string) ([]Question, error)
}

type StoreQuestion struct {
	db *sql.DB
}

func NewStoreQuestion(db *sql.DB) *StoreQuestion {
	return &StoreQuestion{
		db: db,
	}
}
