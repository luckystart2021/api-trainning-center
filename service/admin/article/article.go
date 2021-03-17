package article

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IArticleService interface {
	CountChildArticles(metaChild, metaParent string) (int, error)
	CountArticles(idCategory int) (int, error)
	ShowArticles(idCategory, pageNumber int) ([]Article, error)
	ShowArticlesHomePage(idChildCategory int) ([]Article, error)
	ShowNews() ([]Article, error)
	GetNotificationNews() ([]NotificationNews, error)
	ShowArticlesByChildCategory(idChildCategory int) ([]AdminArticlesList, error)
	ShowArticlesDeleteByChildCategory() ([]AdminArticlesList, error)
	ShowArticlesUnApproval() ([]AdminArticlesList, error)
	GetAllNews() ([]AdminArticlesList, error)
	ShowArticleById(idArticle int) (Articles1, error)
	ShowArticle(idArticle int, meta string) (ArticleDetail, error)
	ShowCategories(idCategoryParent int) ([]CategoriesResponse, error)
	ShowChildArticles(metaChild, metaParent string, pageNo int) ([]ChildCategoryNewsList, error)
	CreateArticle(idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	UpdateArticle(idArticle, idChildCategoryP int, userName, title, description, details, meta, keyWordSEO, image string) (response.MessageResponse, error)
	ApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error)
	UnApprovalArticleById(idArticle int, userName string) (response.MessageResponse, error)
	DeleteArticleById(idArticle int, userName string) (response.MessageResponse, error)
	UnDeleteArticleById(idArticle int, userName string) (response.MessageResponse, error)
	ShowFavoriteNews() ([]ArticlePopular, error)
	ShowResultNewsByKey(searchKey string) ([]Article, error)
}

type StoreArticle struct {
	db *sql.DB
}

func NewStoreArticle(db *sql.DB) *StoreArticle {
	return &StoreArticle{
		db: db,
	}
}
