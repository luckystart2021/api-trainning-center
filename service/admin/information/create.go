package information

import (
	"api-trainning-center/service/response"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (st StoreInformation) CreateInformation(address, phone, email, maps, title, description, img string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
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
		return err
	}
	return nil
}
