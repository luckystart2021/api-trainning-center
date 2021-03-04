package class

import (
	"api-trainning-center/models/admin/class"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreClass) GetListClass() ([]class.Class, error) {
	classLst, err := FindAllClass(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[GetListClass] error : ", err)
		return nil, err
	}
	return classLst, nil
}

func FindAllClass(db *sql.DB) ([]class.Class, error) {
	classLst := []class.Class{}
	query := `
	SELECT
		id,
		code,
		name,
		id_course,
		quantity,
		id_teacher,
		created_by,
		updated_by,
		created_at,
		updated_at,
		is_deleted
	FROM
		class;
	ORDER BY 
		created_at DESC 
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] query error  %v", err)
		return classLst, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	for rows.Next() {
		class := class.Class{}
		var createdAt, updatedAt time.Time
		err = rows.Scan()
		class.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		class.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		classLst = append(classLst, class)
	}

	if len(classLst) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] No Data  %v", err)
		return classLst, errors.New("Không có dữ liệu từ hệ thống")
	}
	return classLst, nil
}
