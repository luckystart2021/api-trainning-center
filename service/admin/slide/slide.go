package slide

import "database/sql"

type ISlideService interface {
	ShowSlides() ([]ShowSlides, error)
}

type StoreSlide struct {
	db *sql.DB
}

func NewStoreSlide(db *sql.DB) *StoreSlide {
	return &StoreSlide{
		db: db,
	}
}
