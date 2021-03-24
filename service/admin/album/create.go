package album

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreAlbum) CreateAlbum(name, meta string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	if err := CreateAlbumByRequest(st.db, name, meta); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Thêm album thành công"
	return response, nil
}

func CreateAlbumByRequest(db *sql.DB, name, meta string) error {
	query := `
	INSERT INTO album (name, meta)
	VALUES($1,$2);
	`
	_, err := db.Exec(query, name, meta)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbumByRequest] Insert album DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
