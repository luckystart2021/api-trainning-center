package album

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/service/response"
	"database/sql"
)

type IAlbumService interface {
	GetListAlbum() ([]photo.Album, error)
	GetAlbumDetail(id int) (photo.AlbumResponse, error)
	CreateAlbum(name, meta string) (response.MessageResponse, error)
	UpdateAlbum(id int, name, meta string) (response.MessageResponse, error)
	DeleteAlbum(id int) (response.MessageResponse, error)
}

type StoreAlbum struct {
	db *sql.DB
}

func NewStoreAlbum(db *sql.DB) *StoreAlbum {
	return &StoreAlbum{
		db: db,
	}
}
