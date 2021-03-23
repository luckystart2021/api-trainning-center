package photo

import (
	"database/sql"
)

type IPhotoService interface {
	ShowPhotos(idAlbum int) ([]PhotosResponse, error)
}

type StorePhoto struct {
	db *sql.DB
}

func NewStorePhoto(db *sql.DB) *StorePhoto {
	return &StorePhoto{
		db: db,
	}
}
