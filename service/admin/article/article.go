package article

import "database/sql"

type IArticleService interface {
	ShowArticles(idCategory int) ([]Article, error)
	ShowArticle(idArticle int, meta string) (ArticleDetail, error)
	ShowCategories(idCategoryParent int) ([]CategoriesResponse, error)
	ShowChildArticles(idChildCategoryP int, meta string) ([]ChildCategoryNewsList, error)
}

type StoreArticle struct {
	db *sql.DB
}

func NewStoreArticle(db *sql.DB) *StoreArticle {
	return &StoreArticle{
		db: db,
	}
}
