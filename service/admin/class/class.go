package class

import (
	"api-trainning-center/models/admin/class"
	"api-trainning-center/service/response"
	"database/sql"
)

type IClassService interface {
	CreateClass(userName, className string, idCource, idTeacher, quantity int64) (response.MessageResponse, error)
	UpdateClass(idClass int, userName, className string, idCource, idTeacher, quantity int64) (response.MessageResponse, error)
	GetListClass() ([]class.Class, error)
}

type StoreClass struct {
	db *sql.DB
}

func NewStoreClass(db *sql.DB) *StoreClass {
	return &StoreClass{
		db: db,
	}
}
