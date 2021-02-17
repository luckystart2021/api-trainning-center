package news

import (
	"api-trainning-center/service/admin/article"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := article.NewStoreArticle(db)
	return func(router chi.Router) {
		router.Get("/{id_category}/news", GetArticles(st))
		router.Get("/{meta}/{id_article}/news", GetArticle(st))
		router.Get("/list/category/{id_category_parent}", GetCategories(st))
		router.Get("/{meta_parent}/{meta}/{id_child_category}/list/child-news", GetChildArticles(st))
	}
}
