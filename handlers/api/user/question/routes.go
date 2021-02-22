package question

import (
	"api-trainning-center/service/admin/question"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := question.NewStoreQuestion(db)
	return func(router chi.Router) {
		// router.Get("/question/{id}", GetQuestionAnswer(st))
		// router.Post("/question/result", GetResult(st))
		// router.Get("/question/exam", GetExam(st))

		router.Get("/question/{rank}/view-suite", ShowSuiteTest(st))
		router.Get("/question/{id_suite}/view-questions", ShowQuestionsByIdSuite(st))
	}
}
