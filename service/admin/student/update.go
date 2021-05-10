package student

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreStudent) UpdateStudent(id int, sex, dayOfBirth, phone, address, fullName, userName string, idClass int, cmnd string, cnsk bool, gplx string, exp, numberKm int, amount float64) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateStudentByRequest(st.db, idClass, sex, dayOfBirth, phone, address, fullName, userName, cmnd, id, cnsk, gplx, exp, numberKm, amount)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateStudentByRequest] Update student DB err  %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin học viên thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin học viên không thành công"
	}
	return resp, nil
}

func updateStudentByRequest(db *sql.DB, idClass int, sex, dayOfBirth, phone, address, fullName, userName, cmnd string, id int, cnsk bool, gplx string, exp, numberKm int, amount float64) (int64, error) {
	timeUpdate := time.Now()
	query := `
	UPDATE
		student
	SET
		sex = $1,
		dateofbirth = $2,
		phone = $3,
		address = $4,
		fullname =$5,
		class_id = $6,
		updated_by = $7,
		updated_at = $8,
		cmnd = $9,
		cnsk = $11,
		gplx =$12,
		experience_driver =$13,
		km_safe =$14,
		amount = $15
	WHERE
		id =$10;
	`
	res, err := db.Exec(query, sex, dayOfBirth, phone, address, fullName, idClass, userName, timeUpdate, cmnd, id, cnsk, gplx, exp, numberKm, amount)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateStudentByRequest] update student in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update student in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
