package student

import (
	"api-trainning-center/models/admin/student"
	"database/sql"
)

type IStudentService interface {
	ShowStudents() ([]student.Student, error)
	ShowStudent(idStudent int) (student.Student, error)
}

type StoreStudent struct {
	db *sql.DB
}

func NewStoreStudent(db *sql.DB) *StoreStudent {
	return &StoreStudent{
		db: db,
	}
}
