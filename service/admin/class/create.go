package class

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreClass) CreateClass(userName, className string, idCource, idTeacher, quantity int64) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := CreateClassByRequest(st.db, userName, className, idCource, idTeacher, quantity); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm lớp thành công"
	return resp, nil
}

func CreateClassByRequest(db *sql.DB, userName, className string, idCource, idTeacher, quantity int64) error {
	query := `
	INSERT INTO class
		(code, name, course_id, quantity, teacher_id, created_by, updated_by) 
	(
	SELECT
		CONCAT('L-', COUNT(*)+1), $1, $2, $3, $4, $5, $6
	FROM
		"class");
	`
	_, err := db.Exec(query, className, idCource, quantity, idTeacher, userName, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateClassByRequest]Insert class DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
