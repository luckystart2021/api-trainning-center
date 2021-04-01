package student

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreStudent) CreateStudent(sex, dayOfBirth, phone, address, fullName, userName string, idClass int, cmnd string, cnsk bool, gplx string, exp, numberKm int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := CreateStudentByRequest(st.db, idClass, sex, dayOfBirth, phone, address, fullName, userName, cmnd, cnsk, gplx, exp, numberKm); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm học viên thành công"
	return resp, nil
}

func CreateStudentByRequest(db *sql.DB, idClass int, sex, dayOfBirth, phone, address, fullName, userName, cmnd string, cnsk bool, gplx string, exp, numberKm int) error {
	query := `
	INSERT INTO student
		(code, sex, dateofbirth, phone, address, fullname, id_class, created_by, updated_by, cmnd, cnsk, gplx, experience_driver, km_safe)
 	(
	SELECT
		CONCAT('HV', COUNT(*)+1), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
	FROM
		student
	);
	`
	_, err := db.Exec(query, sex, dayOfBirth, phone, address, fullName, idClass, userName, userName, cmnd, cnsk, gplx, exp, numberKm)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateStudentByRequest]Insert student DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
