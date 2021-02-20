package article

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IArticleService interface {
	ShowArticles(idChildCategory int) ([]Article, error)
	ShowNews() ([]Article, error)
	ShowArticlesByChildCategory(idChildCategory int) ([]AdminArticlesList, error)
	ShowArticleById(idArticle int) (Articles, error)
	ShowArticle(idArticle int, meta string) (ArticleDetail, error)
	ShowCategories(idCategoryParent int) ([]CategoriesResponse, error)
	ShowChildArticles(metaChild, metaParent string) ([]ChildCategoryNewsList, error)
	CreateArticle(idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	UpdateArticle(idArticle, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	ApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error)
	UnApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error)
	DeleteArticleById(idArticle int, userName string) (response.MessageResponse, error)
	UnDeleteArticleById(idArticle int, userName string) (response.MessageResponse, error)
	ShowFavoriteNews() ([]Article, error)
}

type StoreArticle struct {
	db *sql.DB
}

func NewStoreArticle(db *sql.DB) *StoreArticle {
	return &StoreArticle{
		db: db,
	}
}
