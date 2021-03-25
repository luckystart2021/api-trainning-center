package photo

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/service/response"
	"database/sql"
)

type IPhotoService interface {
	ShowPhotos() ([]PhotosResponse, error)
	ShowPhotosInAdmin() ([]photo.Photo, error)
	ShowPhotoInAdmin(id int) (photo.Photo, error)
	DeletePhoto(id int) (response.MessageResponse, error)
	CreatePhoto(req photo.PhotoRequest, userName string) (response.MessageResponse, error)
	UpdatePhoto(id int, req photo.PhotoRequest, userName string) (response.MessageResponse, error)
}

type StorePhoto struct {
	db *sql.DB
}

func NewStorePhoto(db *sql.DB) *StorePhoto {
	return &StorePhoto{
		db: db,
	}
}
