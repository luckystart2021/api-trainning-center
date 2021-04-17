package subject

import (
	"api-trainning-center/internal/models"
	"database/sql"
)

type ISubjectService interface {
	// CreateSubject(req models.Teacher, userName string) (response.MessageResponse, error)
	// UpdateSubject(id int, req models.Teacher, userName string) (response.MessageResponse, error)
	// ShowSubject(idTeacher int) (models.Teacher, error)
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
