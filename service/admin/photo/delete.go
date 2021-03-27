package photo

import (
	"api-trainning-center/service/response"
	"api-trainning-center/utils"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StorePhoto) DeletePhoto(id int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	photo, err := FindOnePhoto(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeletePhoto] error : ", err)
		return resp, err
	}
	count, err := deletePhotoByRequest(st.db, id)
	if err != nil {
		return resp, err
	}
	err = utils.DeleteFile("upload/img/album/" + photo.Img)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeletePhoto] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa không thành công"
	}
	return resp, nil
}

func deletePhotoByRequest(db *sql.DB, id int) (int64, error) {
	query := `
	DELETE 
		FROM 
			photos
		WHERE 
			id=$1;
	`
	res, err := db.Exec(query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deletePhotoByRequest] delete photos in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] delete photos in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
