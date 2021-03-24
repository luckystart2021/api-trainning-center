package album

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) UpdateAlbum(id int, name, meta string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateAlbumByRequest(st.db, id, name, meta)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateAlbumByRequest] Update album DB err  %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin album thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin album không thành công"
	}
	return resp, nil
}

func updateAlbumByRequest(db *sql.DB, id int, name, meta string) (int64, error) {
	query := `
	UPDATE
		album
	SET
		name = $1,
		meta = $2
	WHERE
		id =$3;
	`
	res, err := db.Exec(query, name, meta, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateAlbumByRequest] update album in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update album in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
