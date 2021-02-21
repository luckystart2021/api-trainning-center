package slide

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreSlide) CreateSlide(userName, title, img string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	if userName == "" {
		userName = "admin1"
	}
	if err := CreateSlideByRequest(st.db, userName, title, img); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Thêm silde thành công"
	return response, nil
}

func CreateSlideByRequest(db *sql.DB, userName, title, img string) error {
	query := `
	INSERT INTO slide (title, img, created_by)
	VALUES($1,$2,$3);
	`
	_, err := db.Exec(query, title, img, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateSlideByRequest]Insertslide DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
