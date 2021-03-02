package class

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IClassService interface {
	CreateClass(userName, className string, idCource, idTeacher, quantity int64) (response.MessageResponse, error)
	UpdateClass(idClass int, userName, className string, idCource, idTeacher, quantity int64) (response.MessageResponse, error)
	// UpdateChildCategory(id int, userName, title, meta string, idCategory int) (response.MessageResponse, error)
	// ShowChildCategories(idCategoryParent int) ([]Categories, error)
	// ShowChildCategory(idChildCategory int) (article.Categories, error)
	// DeleteCategoryById(id int, userName string) (response.MessageResponse, error)
	// UnDeleteCategoryById(id int, userName string) (response.MessageResponse, error)
}

type StoreClass struct {
	db *sql.DB
}

func NewStoreClass(db *sql.DB) *StoreClass {
	return &StoreClass{
		db: db,
	}
}
