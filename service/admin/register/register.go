package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type IRegisterService interface {
	ShowRegisterByClassId(classID int) (models.RegisterGroundSlice, error)
	ShowRegisterById(Id int) (models.RegisterGround, error)
	CreateRegister(req models.RegisterGround) (response.MessageResponse, error)
	UpdateRegister(id int, req models.RegisterGround) (response.MessageResponse, error)
	DeleteRegister(id int) (response.MessageResponse, error)
}

type StoreRegister struct {
	db *sql.DB
}

func NewStoreRegister(db *sql.DB) *StoreRegister {
	return &StoreRegister{
		db: db,
	}
}
