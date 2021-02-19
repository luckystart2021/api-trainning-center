package information

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (tc StoreInformation) ShowInformation() (Information, error) {
	information, err := retrieveInformation(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowInformation] error : ", err)
		return Information{}, err
	}

	return information, nil
}

func retrieveInformation(db *sql.DB) (Information, error) {
	information := Information{}
	query := `
	SELECT
		address, email, phone, maps, title, description, img
	FROM 
		information;
	`
	rows := db.QueryRow(query)
	var img string
	err := rows.Scan(&information.Address, &information.Email, &information.Phone,
		&information.Maps, &information.Title, &information.Description, &img)
	information.Img = "/files/img/information/" + img
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveInformation] Scan error  %v", err)
		return information, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	return information, nil
}
