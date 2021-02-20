package course

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type ICourseService interface {
	ShowCoursesActive() ([]Course, error)
	ShowCourses(idCourse string) (Course, error)
	CreateCourse(userName, name, startDate, endDate, graduationDate, testDate, trainingSystem string) (response.MessageResponse, error)
	UpdateCourse(idCourse int, userName, name, startDate, endDate, graduationDate, testDate, trainingSystem string) (response.MessageResponse, error)
}

type StoreCourse struct {
	db *sql.DB
}

func NewStoreCourse(db *sql.DB) *StoreCourse {
	return &StoreCourse{
		db: db,
	}
}
