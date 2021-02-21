package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

const (
	isHide   = true
	isUnHide = false
)

func (tc StoreSlide) HideSlideById(idSlide int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := hideSlideById(tc.db, idSlide, isHide)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[HideSlideById]Hide slide DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Ẩn slide thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy slide"
	}
	return resp, nil
}

func (tc StoreSlide) UnHideSlideById(idSlide int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := hideSlideById(tc.db, idSlide, isUnHide)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UnHideSlideById]UnHide slide DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Hiện thị slide thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy slide"
	}
	return resp, nil
}

func hideSlideById(db *sql.DB, id int, isHide bool) (int64, error) {
	query := `
	UPDATE
		slide
	SET
		hide = $2
	WHERE
		id = $1
	`
	res, err := db.Exec(query, id, isHide)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[hideSlideById] hide slide DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] hide slide DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
