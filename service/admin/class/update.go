package class

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreClass) UpdateClass(idClass int, userName, className string, idCource, idTeacher, quantity int64, isDeleted bool) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := UpdateClassByRequest(st.db, idClass, userName, className, idCource, idTeacher, quantity, isDeleted); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Cập nhật lớp thành công"
	return resp, nil
}

func UpdateClassByRequest(db *sql.DB, idClass int, userName, className string, idCource, idTeacher, quantity int64, isDeleted bool) error {
	timeUpdate := time.Now()
	query := `
	UPDATE
		class
	SET
		"name" = $1,
		course_id = $2,
		quantity = $3,
		teacher_id = $4,
		updated_by = $5,
		updated_at = $6,
		is_deleted = $8
	WHERE
		id = $7;
	`
	_, err := db.Exec(query, className, idCource, quantity, idTeacher, userName, timeUpdate, idClass, isDeleted)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateClassByRequest] update class DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
