package schedule

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IScheduleService interface {
	GenerateSchedule(courseID int) (response.MessageResponse, error)
}

type StoreSchedule struct {
	db *sql.DB
}

func NewStoreSchedule(db *sql.DB) *StoreSchedule {
	return &StoreSchedule{
		db: db,
	}
}
