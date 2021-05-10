package student

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreStudent) CreateStudent(sex, dayOfBirth, phone, address, fullName, userName string, idClass int, cmnd string, cnsk bool, gplx string, exp, numberKm int, amount float64) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()
	countStudent, err := models.Students(
		qm.Where("class_id = ?", idClass),
	).Count(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CountStudent]Count student DB err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	class, err := models.FindClass(ctx, st.db, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CountClass]Count class DB err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	if countStudent >= int64(class.Quantity) {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateStudent]Create Student DB err  %v", err)
		return resp, errors.New("Số lượng học viên vượt quá danh sách lớp")
	}

	if err := CreateStudentByRequest(st.db, idClass, sex, dayOfBirth, phone, address, fullName, userName, cmnd, cnsk, gplx, exp, numberKm, amount); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm học viên thành công"
	return resp, nil
}

func CreateStudentByRequest(db *sql.DB, idClass int, sex, dayOfBirth, phone, address, fullName, userName, cmnd string, cnsk bool, gplx string, exp, numberKm int, amount float64) error {
	query := `
	INSERT INTO student
		(code, sex, dateofbirth, phone, address, fullname, class_id, created_by, updated_by, cmnd, cnsk, gplx, experience_driver, km_safe, amount)
 	(
	SELECT
		CONCAT('HV', COUNT(*)+1), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
	FROM
		student
	);
	`
	_, err := db.Exec(query, sex, dayOfBirth, phone, address, fullName, idClass, userName, userName, cmnd, cnsk, gplx, exp, numberKm, amount)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateStudentByRequest]Insert student DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
