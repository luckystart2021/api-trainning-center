package slide

import (
	"api-trainning-center/models/admin/slide"
	"api-trainning-center/service/response"
	"database/sql"
)

type ISlideService interface {
	ShowSlidesAdmin() ([]slide.Slide, error)
	ShowDetailSlide(idSlide int) (slide.Slide, error)
	ShowSlides() ([]ShowSlides, error)
	CreateSlide(userName, title, img string) (response.MessageResponse, error)
	UpdateSlide(id int, title, img string) (response.MessageResponse, error)
	HideSlideById(idSlide int) (response.MessageResponse, error)
}

type StoreSlide struct {
	db *sql.DB
}

func NewStoreSlide(db *sql.DB) *StoreSlide {
	return &StoreSlide{
		db: db,
	}
}
