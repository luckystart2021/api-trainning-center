package question

import (
	"api-trainning-center/service/admin/question"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := question.NewStoreQuestion(db)
	return func(router chi.Router) {
		router.Get("/question/rank", ShowRankVehicle(st))
		router.Get("/question/{id_rank}/view-suite", ShowSuiteTest(st))
		router.Get("/question/{id_suite}/view-questions", ShowQuestionsByIdSuite(st))
		router.Post("/question/result", GetResult(st))
	}
}
