package schedule

import (
	"database/sql"
)

type IScheduleService interface {
	GenerateSchedule(courseID int) ([]ScheduleResponse, error)
}

type StoreSchedule struct {
	db *sql.DB
}

func NewStoreSchedule(db *sql.DB) *StoreSchedule {
	return &StoreSchedule{
		db: db,
	}
}
