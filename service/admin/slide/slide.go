package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type ISlideService interface {
	ShowSlides() ([]ShowSlides, error)
	CreateSlide(userName, title, img string) (response.MessageResponse, error)
}

type StoreSlide struct {
	db *sql.DB
}

func NewStoreSlide(db *sql.DB) *StoreSlide {
	return &StoreSlide{
		db: db,
	}
}
