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

func (st StoreClass) GetDetailClass(idClass int) (class.Class, error) {
	cl := class.Class{}
	class, err := FindOneClass(st.db, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindOneClass] error : ", err)
		return cl, err
	}
	return class, nil
}

func FindOneClass(db *sql.DB, idStudent int) (class.Class, error) {
	class := class.Class{}
	query := `
	SELECT
		id,
		code,
		name,
		id_course,
		quantity,
		id_teacher,
		is_deleted,
		created_by,
		updated_by,
		created_at,
		updated_at
	FROM
		class
	WHERE
		id = $1
	ORDER BY 
		created_at DESC;
	`
	rows := db.QueryRow(query, idStudent)
	var createdAt, updatedAt time.Time
	err := rows.Scan(&class.Id, &class.Code, &class.Name, &class.IdCourse, &class.Quantity, &class.IdTeacher, &class.IsDelete, &class.CreatedBy, &class.UpdatedBy, &createdAt, &updatedAt)
	class.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	class.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneClass] No Data  %v", err)
		return class, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneClass] Scan error  %v", err)
	}
	return class, nil
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
		is_deleted,
		created_by,
		updated_by,
		created_at,
		updated_at
	FROM
		class
	ORDER BY 
		created_at DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] query error  %v", err)
		return classLst, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		class := class.Class{}
		var createdAt, updatedAt time.Time
		err = rows.Scan(&class.Id, &class.Code, &class.Name, &class.IdCourse, &class.Quantity, &class.IdTeacher, &class.IsDelete, &class.CreatedBy, &class.UpdatedBy, &createdAt, &updatedAt)
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
