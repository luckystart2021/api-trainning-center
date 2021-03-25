package photo

import (
	"api-trainning-center/models/admin/photo"
	"database/sql"
)

type IPhotoService interface {
	ShowPhotos(idAlbum int) ([]PhotosResponse, error)
	ShowPhotosInAdmin() ([]photo.Photo, error)
	ShowPhotoInAdmin(id int) (photo.Photo, error)
}

type StorePhoto struct {
	db *sql.DB
}

func NewStorePhoto(db *sql.DB) *StorePhoto {
	return &StorePhoto{
		db: db,
	}
}
