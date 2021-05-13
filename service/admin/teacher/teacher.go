package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type ITeacherService interface {
	CreateTeacher(req models.Teacher, userName string) (response.MessageResponse, error)
	UpdateTeacher(id int, req models.Teacher, userName string) (response.MessageResponse, error)
	ShowTeachers() (models.TeacherSlice, error)
	ShowTeacher(idTeacher int) (models.Teacher, error)
	ShowTeacherByAvalible() (models.TeacherSlice, error)
	InActive(idTeacher int, userName string) (response.MessageResponse, error)
	Active(idTeacher int, userName string) (response.MessageResponse, error)
}

type StoreTeacher struct {
	db *sql.DB
}

func NewStoreTeacher(db *sql.DB) *StoreTeacher {
	return &StoreTeacher{
		db: db,
	}
}
