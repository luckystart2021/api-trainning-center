package information

import (
	"api-trainning-center/models/admin/information"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreInformation) CreateInformation(address, phone, email, maps, title, description, img string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	count, err := information.CountInformation(st.db)
	if err != nil {
		return response, err
	}
	if count >= 1 {
		return response, errors.New("Thông tin website đã tồn tại")
	}
	if err := CreateInformationByRequest(st.db, address, phone, email, maps, title, description, img); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Thêm thông tin thành công"
	return response, nil
}

func CreateInformationByRequest(db *sql.DB, address, phone, email, maps, title, description, img string) error {
	query := `
	INSERT INTO information
		(address, email, phone, maps, title, description, img)
	VALUES($1, $2, $3, $4, $5, $6, $7);
	`
	_, err := db.Exec(query, address, phone, email, maps, title, description, img)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateInformationByRequest]Insert contact DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
