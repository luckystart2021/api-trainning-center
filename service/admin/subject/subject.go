package subject

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type ISubjectService interface {
	CreateSubject(req models.Subject, userName string) (response.MessageResponse, error)
	UpdateSubject(subjectID int, req models.Subject, userName string) (response.MessageResponse, error)
	ShowSubject(subjectID int) (models.Subject, error)
	ShowSubjects() (models.SubjectSlice, error)
}

type StoreSubject struct {
	db *sql.DB
}

func NewStoreSubject(db *sql.DB) *StoreSubject {
	return &StoreSubject{
		db: db,
	}
}
