package course

import (
	"database/sql"
)

type ICourseService interface {
	ShowCoursesActive() ([]Course, error)
}

type StoreCourse struct {
	db *sql.DB
}

func NewStoreCourse(db *sql.DB) *StoreCourse {
	return &StoreCourse{
		db: db,
	}
}
