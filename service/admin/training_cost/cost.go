package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type ICostService interface {
	// CreateTeacher(req models.Teacher, userName string) (response.MessageResponse, error)
	// UpdateTeacher(id int, req models.Teacher, userName string) (response.MessageResponse, error)
	// ShowTeachers() (models.TeacherSlice, error)
	// ShowTeacher(idTeacher int) (models.Teacher, error)
	// ShowTeacherByAvalible() (models.TeacherSlice, error)
	// InActive(idTeacher int, userName string) (response.MessageResponse, error)
	// Active(idTeacher int, userName string) (response.MessageResponse, error)
	CreateCost(req models.TrainingCost, courseID, classID int, userName string) (response.MessageResponse, error)
	UpdateCost(id int, req models.TrainingCost, userName string) (response.MessageResponse, error)
	ShowCost(courseID int) ([]TrainingCost, error)
	ShowDetailCost(costID int) (TrainingCost, error)
}

type StoreCost struct {
	db *sql.DB
}

func NewStoreCost(db *sql.DB) *StoreCost {
	return &StoreCost{
		db: db,
	}
}
