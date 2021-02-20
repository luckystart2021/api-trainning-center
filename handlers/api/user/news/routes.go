package news

import (
	"api-trainning-center/service/admin/article"
	"database/sql"

	"github.com/go-chi/chi"
)

func Router(db *sql.DB) func(chi.Router) {
	st := article.NewStoreArticle(db)
	return func(router chi.Router) {
		router.Get("/{id_category}/news/pagination", GetArticlePagination(st))
		router.Get("/{id_category}/news", GetArticles(st))
		router.Get("/{meta_article}/{id_article}/news", GetArticle(st))
		router.Get("/list/category/{id_category_parent}", GetCategories(st))
		router.Get("/{meta_parent}/{meta_child}/list/child-news", GetChildArticles(st))
		router.Get("/{meta_parent}/{meta_child}/list/child-news/pagination", GetChildArticlesPagination(st))
		router.Get("/news", GetNews(st))
		router.Get("/popular/news", GetFavoriteNews(st))
		router.Get("/news/search", SearchNews(st))
	}
}
