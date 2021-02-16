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
	}
}
