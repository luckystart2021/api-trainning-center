package class

import (
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreClass) UpdateClass(idClass int, userName string, idCource, quantity int64, isDeleted bool) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := UpdateClassByRequest(st.db, idClass, userName, idCource, quantity, isDeleted); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Cập nhật lớp thành công"
	return resp, nil
}

func UpdateClassByRequest(db *sql.DB, idClass int, userName string, idCource, quantity int64, isDeleted bool) error {
	timeUpdate := time.Now()
	query := `
	UPDATE
		class
	SET
		course_id = $1,
		quantity = $2,
		updated_by = $3,
		updated_at = $4,
		is_deleted = $6
	WHERE
		id = $5;
	`
	_, err := db.Exec(query, idCource, quantity, userName, timeUpdate, idClass, isDeleted)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateClassByRequest] update class DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
