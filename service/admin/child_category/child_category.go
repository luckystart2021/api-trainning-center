package child_category

import (
	"api-trainning-center/service/admin/article"
	"api-trainning-center/service/response"
	"database/sql"
)

type IChildCategoryService interface {
	CreateChildCategory(userName, title, meta string, idCategory int) (response.MessageResponse, error)
	UpdateChildCategory(id int, userName, title, meta string, idCategory int) (response.MessageResponse, error)
	ShowChildCategories(idCategoryParent int) ([]article.Categories, error)
}

type StoreChildCategory struct {
	db *sql.DB
}

func NewStoreChildCategory(db *sql.DB) *StoreChildCategory {
	return &StoreChildCategory{
		db: db,
	}
}
