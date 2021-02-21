package information

import (
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

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

func (tc StoreInformation) ShowInformationAdmin() (InformationAdmin, error) {
	information, err := retrieveInformationAdmin(tc.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowInformationAdmin] error : ", err)
		return InformationAdmin{}, err
	}

	return information, nil
}

func retrieveInformationAdmin(db *sql.DB) (InformationAdmin, error) {
	information := InformationAdmin{}
	query := `
	SELECT
		id,address, email, phone, maps, title, description, img, created_at
	FROM 
		information;
	`
	rows := db.QueryRow(query)
	var img string
	var createdAt time.Time
	err := rows.Scan(&information.Id, &information.Address, &information.Email, &information.Phone,
		&information.Maps, &information.Title, &information.Description, &img, &createdAt)
	information.Img = "/files/img/information/" + img
	information.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveInformationAdmin] No Data  %v", err)
		return information, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[retrieveInformationAdmin] Scan error  %v", err)
		return information, errors.New("Lỗi hệ thống vui lòng thử lại")
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
