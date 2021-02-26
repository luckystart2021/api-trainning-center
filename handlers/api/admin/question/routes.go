package question

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/question"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := question.NewStoreQuestion(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Post("/question/create", CreateQuestion(st))
		router.Get("/questions/view", ShowQuestions(st))
		router.Get("/question/{id_question}/view", ShowQuestion(st))
		router.Put("/question/{id_question}/update", UpdateQuestion(st))
		router.Delete("/question/delete/{id_question}", DeleteQuestion(st))
	}
}
