package article

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IArticleService interface {
	ShowArticles(idChildCategory int) ([]Article, error)
	ShowArticlesByChildCategory(idChildCategory int) ([]AdminArticlesList, error)
	ShowArticleById(idArticle int) (Articles, error)
	ShowArticle(idArticle int, meta string) (ArticleDetail, error)
	ShowCategories(idCategoryParent int) ([]CategoriesResponse, error)
	ShowChildArticles(idChildCategoryP int, meta string) ([]ChildCategoryNewsList, error)
	CreateArticle(idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	UpdateArticle(idArticle, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	DeleteArticleById(idArticle int) (response.MessageResponse, error)
}

type StoreArticle struct {
	db *sql.DB
}

func NewStoreArticle(db *sql.DB) *StoreArticle {
	return &StoreArticle{
		db: db,
	}
}
