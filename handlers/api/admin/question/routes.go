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
		router.Get("/questions/{code_de}/view", ShowQuestions(st))
	}
}
