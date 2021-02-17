package news

import (
	"api-trainning-center/middlewares"
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/constant"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func Router(db *sql.DB, client *redis.Client) func(chi.Router) {
	st := article.NewStoreArticle(db)
	return func(router chi.Router) {
		router.Use(middlewares.AuthJwtVerify)
		router.Use(middlewares.CheckScopeAccess(client, constant.ADMIN))
		router.Post("/article/create", CreateArticle(st))
		router.Put("/article/{id_article}/update", UpdateArticle(st))
		router.Get("/article/{id_child_article}/views", ShowArticles(st))
		router.Get("/article/{id_article}/detail", ShowDetailArticle(st))
		router.Put("/article/{id_article}/delete", DeleteArticle(st))
		router.Put("/article/{id_article}/approval", ApprovalArticle(st))
	}
}
