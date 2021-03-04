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
			router.Get("/{id_child_category}/views", ShowArticles(st))
			router.Post("/create", CreateArticle(st))
			router.Get("/views-deleted", ShowArticlesDeleted(st))
			router.Get("/views-un-approval", ShowArticlesUnApproval(st))
			router.Get("/views-all", ShowAllNews(st))
			router.Route("/{id_article}", func(router chi.Router) {
				router.Put("/update", UpdateArticle(st))
				router.Get("/detail", ShowDetailArticle(st))
				router.Put("/approval", ApprovalArticle(st))
				router.Put("/un-approval", UnApprovalArticle(st))
				router.Put("/delete", DeleteArticle(st))
				router.Put("/un-delete", UnDeleteArticle(st))
			})
		})
	}
}
