package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSlide) UpdateSlide(id int, title, img string) (response.MessageResponse, error) {
	response := response.MessageResponse{}

	if err := updateSlideByRequest(st.db, id, title, img); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Cập nhật slide thành công"
	return response, nil
}

func updateSlideByRequest(db *sql.DB, id int, title, img string) error {
	if img == "" || len(img) == 0 {
		query := `
		UPDATE
			slide
		SET
			title = $1
		WHERE
			id = $2;
		`
		_, err := db.Exec(query, title, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateSlideByRequest] update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	} else {
		query := `
		UPDATE
			slide
		SET
			title = $1,
			img = $2
		WHERE
			id = $3;
		`
		_, err := db.Exec(query, title, img, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[updateSlideByRequest] update DB err  %v", err)
			return errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
	}

	return nil
}
