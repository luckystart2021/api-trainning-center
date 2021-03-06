package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSlide) UpdateSlide(id int, title, img, hide string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateSlideByRequest(st.db, id, title, img, hide)
	if err != nil {
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật slide thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật slide không thành công"
	}
	return resp, nil
}

func updateSlideByRequest(db *sql.DB, id int, title, img, hide string) (int64, error) {
	var rowsAffected int64
	if img == "" || len(img) == 0 {
		query := `
		UPDATE
			slide
		SET
			title = $1,
			hide = $3
		WHERE
			id = $2;
		`
		res, err := db.Exec(query, title, id, hide)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateSlideByRequest] update DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	} else {
		query := `
		UPDATE
			slide
		SET
			title = $1,
			img = $2,
			hide = $4
		WHERE
			id = $3;
		`
		res, err := db.Exec(query, title, img, id, hide)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateSlideByRequest] update DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
		// check how many rows affected
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateArticleByRequest] Update Article DB err  %v", err)
			return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	return rowsAffected, nil
}
