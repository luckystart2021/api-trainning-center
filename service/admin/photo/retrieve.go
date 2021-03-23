package photo

import (
	"api-trainning-center/models/admin/photo"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

type PhotoResponse struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Meta  string `json:"meta"`
}

type PhotosResponse struct {
	AlbumName string          `json:"album_name"`
	AlbumMeta string          `json:"album_meta"`
	Photos    []PhotoResponse `json:"photo"`
}

func (st StorePhoto) ShowPhotos(idAlbum int) ([]PhotosResponse, error) {
	album, err := findAlbum(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAlbum] error : ", err)
		return nil, err
	}
	photosResp := []PhotosResponse{}
	for _, dataAlbum := range album {
		photosResponse := PhotosResponse{}
		photosResponse.AlbumName = dataAlbum.Name
		photosResponse.AlbumMeta = dataAlbum.Meta
		photos, err := findPhotosByIdAlbum(st.db, dataAlbum.Id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[findPhotosByIdAlbum] error : ", err)
			return nil, err
		}
		photosR := []PhotoResponse{}
		for _, data := range photos {
			photo := PhotoResponse{}
			photo.Img = "/files/img/album/" + data.Img
			photo.Meta = data.Meta
			photo.Title = data.Title
			photosR = append(photosR, photo)
		}
		photosResponse.Photos = photosR
		photosResp = append(photosResp, photosResponse)
	}

	return photosResp, nil
}

func findAlbum(db *sql.DB) ([]photo.Album, error) {
	albums := []photo.Album{}
	query := `
	SELECT
		id,
		name,
		meta
	FROM
		album
	ORDER BY id;
	`
	rows, err := db.Query(query)

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] query error  %v", err)
		return albums, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		album := photo.Album{}
		err = rows.Scan(&album.Id, &album.Name, &album.Meta)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] Scan error  %v", err)
			return albums, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		albums = append(albums, album)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(albums) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAlbum] No Data  %v", err)
		return albums, errors.New("Không có dữ liệu từ hệ thống")
	}
	return albums, nil
}

func findPhotosByIdAlbum(db *sql.DB, idAlbum int) ([]photo.Photo, error) {
	photos := []photo.Photo{}
	query := `
	SELECT
		p.id,
		img,
		title,
		p.meta,
		created_by,
		created_at,
		updated_by,
		updated_at,
		id_album
	FROM
		photos p
	INNER JOIN 
		album a ON a.id = p.id_album
	WHERE
		id_album = $1;
	`
	rows, err := db.Query(query, idAlbum)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] query error  %v", err)
		return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	defer rows.Close()
	for rows.Next() {
		photo := photo.Photo{}
		err = rows.Scan(&photo.Id, &photo.Img, &photo.Title, &photo.Meta, &photo.CreatedBy, &photo.CreatedAt, &photo.UpdatedBy, &photo.UpdatedAt, &photo.IdAlbum)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] Scan error  %v", err)
			return photos, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		photos = append(photos, photo)
	}

	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findPhotosByIdAlbum] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return photos, nil
}
