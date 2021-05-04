package child_subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type IChildSubjectService interface {
	CreateChildSubject(req models.ChildSubject, userName string) (response.MessageResponse, error)
	UpdateChildSubject(childSubjectID int, req models.ChildSubject, userName string) (response.MessageResponse, error)
	ShowChildSubject(childSubjectID int) (models.ChildSubject, error)
	ShowChildSubjects(idChildSubject int) (models.ChildSubjectSlice, error)
	DeleteChildSubject(childSubjectID int) (response.MessageResponse, error)
}

type StoreChildSubject struct {
	db *sql.DB
}

func NewStoreChildSubject(db *sql.DB) *StoreChildSubject {
	return &StoreChildSubject{
		db: db,
	}
}
