package information

import (
	"api-trainning-center/service/response"
	"database/sql"

	"github.com/sirupsen/logrus"
)

func (st StoreInformation) UpdateInformation(idInformationI int, address, phone, email, maps, title, description, img string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	if err := updateInformationByRequest(st.db, idInformationI, address, phone, email, maps, title, description, img); err != nil {
		return response, err
	}
	response.Status = true
	response.Message = "Cập nhật thông tin thành công"
	return response, nil
}

func updateInformationByRequest(db *sql.DB, idInformationI int, address, phone, email, maps, title, description, img string) error {
	query := `
	UPDATE information SET 
		address=$1, 
		email=$2, 
		phone=$3,
		maps=$4,
		title=$5,
		description=$6,
		img=$7
	WHERE id= $8;
	`
	_, err := db.Exec(query, address, phone, email, maps, title, description, img, idInformationI)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateInformationByRequest] Update DB err  %v", err)
		return err
	}
	return nil
}
