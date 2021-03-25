package photo

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StorePhoto) UpdatePhoto(id int, req photo.PhotoRequest, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updatePhotoByRequest(st.db, id, req, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updatePhotoByRequest] Update photo DB err  %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật hình ảnh thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật hình ảnh không thành công"
	}
	return resp, nil
}

func updatePhotoByRequest(db *sql.DB, id int, req photo.PhotoRequest, userName string) (int64, error) {
	timeUpdate := time.Now()
	var rowsAffected int64
	if req.Img == "" {
		query := `
		UPDATE
			photos
		SET
			id_album = $1,
			title = $2,
			meta = $3,
			updated_by = $4,
			updated_at = $5
		WHERE
			id = $6;
		`
		res, err := db.Exec(query, req.IdAlbum, req.Title, req.Meta, userName, timeUpdate, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updatePhotoByRequest] update photo in DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update photo in DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	} else {
		query := `
		UPDATE
			photos
		SET
			id_album = $1,
			img = $2,
			title = $3,
			meta = $4,
			updated_by = $5,
			updated_at = $6
		WHERE
			id = $7;
		`
		res, err := db.Exec(query, req.IdAlbum, req.Img, req.Title, req.Meta, userName, timeUpdate, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updatePhotoByRequest] update photo in DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update photo in DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	return rowsAffected, nil
}
