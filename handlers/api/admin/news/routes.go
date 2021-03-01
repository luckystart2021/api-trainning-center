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
		router.Route("/article", func(router chi.Router) {
			router.Post("/create", CreateArticle(st))
			router.Get("/{id_child_category}/views", ShowArticles(st))
			router.Get("/views-deleted", ShowArticlesDeleted(st))
			router.Get("/views-un-approval", ShowArticlesUnApproval(st))
			router.Put("/{id_article}/update", UpdateArticle(st))
			router.Get("/{id_article}/detail", ShowDetailArticle(st))
			router.Put("/{id_article}/approval", ApprovalArticle(st))
			router.Put("/{id_article}/un-approval", UnApprovalArticle(st))
			router.Put("/{id_article}/delete", DeleteArticle(st))
			router.Put("/{id_article}/un-delete", UnDeleteArticle(st))
		})
	}
}
