package holiday

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type IHolidayService interface {
	CreateHoliday(req models.Holiday) (response.MessageResponse, error)
	UpdateHoliday(idHoliday int, req models.Holiday) (response.MessageResponse, error)
	ShowHolidays() (models.HolidaySlice, error)
	ShowHoliday(idHoliday int) (models.Holiday, error)
	DeleteHoliday(idHoliday int) (response.MessageResponse, error)
}

type StoreHoliday struct {
	db *sql.DB
}

func NewStoreHoliday(db *sql.DB) *StoreHoliday {
	return &StoreHoliday{
		db: db,
	}
}
