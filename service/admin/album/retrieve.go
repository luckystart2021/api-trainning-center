package album

import (
	"api-trainning-center/models/admin/photo"
	photoService "api-trainning-center/service/admin/photo"
	"api-trainning-center/utils"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) GetListAlbum() ([]photo.Album, error) {
	albumLst, err := photoService.FindAlbum(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindAlbum] error : ", err)
		return nil, err
	}
	return albumLst, nil
}

func (st StoreAlbum) GetAlbumDetail(id int) (photo.AlbumResponse, error) {
	photoR := photo.AlbumResponse{}
	album, err := FindOneAlbum(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindOneAlbum] error : ", err)
		return photo.AlbumResponse{}, err
	}
	photoR.Id = album.Id
	photoR.Meta = album.Meta
	photoR.Name = album.Name
	photos, err := photoService.FindPhotosByIdAlbum(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindPhotosByIdAlbum] error : ", err)
		return photo.AlbumResponse{}, err
	}
	photosR := []photo.PhotoResponse{}
	for _, data := range photos {
		photo := photo.PhotoResponse{}
		photo.Id = data.Id
		photo.IdAlbum = data.IdAlbum
		photo.Img = "/files/img/album/" + data.Img
		photo.Meta = data.Meta
		photo.Title = data.Title
		photo.CreatedAt = utils.TimeIn(data.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		photo.CreatedBy = data.CreatedBy
		photo.UpdatedAt = utils.TimeIn(data.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		photo.UpdatedBy = data.UpdatedBy
		photosR = append(photosR, photo)
	}
	photoR.Photos = photosR
	return photoR, nil
}

func FindOneAlbum(db *sql.DB, idAlbum int) (photo.Album, error) {
	album := photo.Album{}
	query := `
	SELECT
		id,
		name,
		meta
	FROM
		album
	WHERE
		id = $1;
	`
	rows := db.QueryRow(query, idAlbum)
	err := rows.Scan(&album.Id, &album.Name, &album.Meta)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneAlbum] No Data  %v", err)
		return album, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneAlbum] Scan error  %v", err)
	}
	return album, nil
}
