package album

import (
	"api-trainning-center/models/admin/photo"
	photoService "api-trainning-center/service/admin/photo"
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

func (st StoreAlbum) GetAlbumDetail(id int) (photo.Album, error) {
	album, err := FindOneAlbum(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindOneAlbum] error : ", err)
		return photo.Album{}, err
	}
	return album, nil
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
