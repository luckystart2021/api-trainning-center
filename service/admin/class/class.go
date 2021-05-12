package class

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/models/admin/class"
	"api-trainning-center/service/response"
	"database/sql"
)

type IClassService interface {
	CreateClass(userName string, idCource, quantity, teacherId, vehicleId int64) (response.MessageResponse, error)
	UpdateClass(idClass int, userName string, idCource, quantity int64, isDeleted bool, teacherId, vehicleId int64) (response.MessageResponse, error)
	GetListClass(idCourse int) (models.ClassSlice, error)
	GetDetailClass(idClass int) (class.Class, error)
	InActiveClass(idClass int, userName string) (response.MessageResponse, error)
	ActiveClass(idClass int, userName string) (response.MessageResponse, error)
}

type StoreClass struct {
	db *sql.DB
}

func NewStoreClass(db *sql.DB) *StoreClass {
	return &StoreClass{
		db: db,
	}
}
