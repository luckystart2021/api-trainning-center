package suite

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/suite"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := suite.NewStoreSuiteTest(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Get("/genarate/{number_suite}/{rank}", GenarateTest(st))
	}
}
