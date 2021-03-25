package photo

import (
	"api-trainning-center/models/admin/photo"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StorePhoto) CreatePhoto(req photo.PhotoRequest, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := createPhotoByRequest(st.db, req, userName); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm ảnh thành công"
	return resp, nil
}

func createPhotoByRequest(db *sql.DB, req photo.PhotoRequest, userName string) error {
	query := `
	INSERT INTO photos
		(img, id_album, title, meta, created_by, updated_by)
	VALUES($1,$2,$3,$4,$5,$6);
	`
	_, err := db.Exec(query, req.Img, req.IdAlbum, req.Title, req.Meta, userName, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[createPhotoByRequest]Insert photo DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
