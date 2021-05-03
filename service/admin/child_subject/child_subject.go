package child_subject

import (
	"api-trainning-center/internal/models"
	"database/sql"
)

type IChildSubjectService interface {
	// CreateSubject(req models.Subject, userName string) (response.MessageResponse, error)
	// UpdateSubject(subjectID int, req models.Subject, userName string) (response.MessageResponse, error)
	ShowChildSubject(childSubjectID int) (models.ChildSubject, error)
	ShowChildSubjects(idChildSubject int) (models.ChildSubjectSlice, error)
}

type StoreChildSubject struct {
	db *sql.DB
}

func NewStoreChildSubject(db *sql.DB) *StoreChildSubject {
	return &StoreChildSubject{
		db: db,
	}
}
