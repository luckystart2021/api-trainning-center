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
	IdQuestion    int    `json:"id_question"`
	IdAnswer      string `json:"answer"`
	CorrectAnswer string `json:"correct_answer"`
}

type QuestionSystem struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	AnswerA       string `json:"answer_a"`
	AnswerB       string `json:"answer_b"`
	AnswerC       string `json:"answer_c"`
	AnswerD       string `json:"answer_d"`
	Img           string `json:"img"`
	AnwserCorrect string `json:"anwser_correct"`
	Liet          bool   `json:"liet"`
}

type ResponseQuestionExam struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//V1

type TestSuiteResponse struct {
	Time  int         `json:"time"`
	Suite []TestSuite `json:"suite"`
}
type TestSuite struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type QuestionResponse struct {
	Id           int       `json:"id"`
	QuestionName string    `json:"question_name"`
	Img          string    `json:"img"`
	Answers      []Answers `json:"answers"`
}

type Questions struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AnswerA string `json:"answer_a"`
	AnswerB string `json:"answer_b"`
	AnswerC string `json:"answer_c"`
	AnswerD string `json:"answer_d"`
	Img     string `json:"img"`
}
type IQuestionService interface {
	ShowQuestions(code string) ([]Question, error)
	ShowQuestionsExam() ([]ResponseQuestionExam, error)
	DeleteQuestion(idQuestion string) (response.MessageResponse, error)
	// V1
	CreateQuestion(name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error)
	GetAllTestSuiteByRank(rank int) (TestSuiteResponse, error)
	GetQuestionsByIdSuite(idSuite int) ([]QuestionResponse, error)
	GetAllRankVehicle() ([]RankResponse, error)
	ShowResult(result result.Result) (ResponseResult, error)
	ShowQuestionsSystem() ([]QuestionSystem, error)
	ShowQuestionSystem(idQuestion string) (QuestionSystem, error)
	UpdateQuestion(idQuestion int, name, answerA, answerB, answerC, answerD, img, result string, liet bool) (response.MessageResponse, error)
}

type StoreQuestion struct {
	db *sql.DB
}

func NewStoreQuestion(db *sql.DB) *StoreQuestion {
	return &StoreQuestion{
		db: db,
	}
}
