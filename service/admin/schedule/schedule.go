package schedule

import (
	"database/sql"
)

type IScheduleService interface {
	GenerateSchedule(courseID int) (Schedule, error)
	RetrieveSchedule(courseID int) (Schedule, error)
	SearchScheduleByCourseCode(code string) (Schedule, error)
}

type StoreSchedule struct {
	db *sql.DB
}

func NewStoreSchedule(db *sql.DB) *StoreSchedule {
	return &StoreSchedule{
		db: db,
	}
}
