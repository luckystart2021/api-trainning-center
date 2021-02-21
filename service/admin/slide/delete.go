package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreSlide) DeleteSlideById(idSlide int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := deleteSlideById(tc.db, idSlide)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteSlideById]delete slide DB err  %v", err)
		return resp, err
	}

	if count > 0 {
		resp.Status = true
		resp.Message = "Xóa slide thành công"
	} else {
		resp.Status = false
		resp.Message = "Không tìm thấy slide"
	}
	return resp, nil
}

func deleteSlideById(db *sql.DB, id int) (int64, error) {
	query := `
	DELETE FROM slide
	WHERE id=$1;
	`
	res, err := db.Exec(query, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteSlideById] delete slide DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[deleteSlideById] delete slide DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
