package contact

import (
	"api-trainning-center/service/response"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (st StoreContact) CreateContact(fullName, phone, email, message, subject string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	if err := CreateContactByRequest(st.db, fullName, phone, email, message, subject); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Gửi thông tin liên hệ thành công"
	return response, nil
}

func CreateContactByRequest(db *sql.DB, fullName, phone, email, message, subject string) error {
	query := `
	INSERT INTO contact
		(fullname, phone, email, message, subject)
	VALUES($1, $2, $3, $4, $5);
	`
	_, err := db.Exec(query, fullName, phone, email, message, subject)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateContactByRequest]Insert contact DB err  %v", err)
		return err
	}
	return nil
}
