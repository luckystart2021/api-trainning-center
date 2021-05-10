package class

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/models/admin/class"
	"api-trainning-center/utils"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// func (st StoreClass) GetListClass(idCourse int) ([]class.Class, error) {
// 	classLst, err := FindAllClass(st.db)
// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{}).Error("[GetListClass] error : ", err)
// 		return nil, err
// 	}
// 	return classLst, nil
// }

func (st StoreClass) GetListClass(idCourse int) (models.ClassSlice, error) {
	ctx := context.Background()
	classes, err := models.Classes(
		qm.Where("course_id = ?", idCourse),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAllClasses] error : ", err)
		return nil, err
	}
	if classes == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return classes, nil
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
		course_id,
		quantity,
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
	err := rows.Scan(&class.Id, &class.Code, &class.IdCourse, &class.Quantity, &class.IsDelete, &class.CreatedBy, &class.UpdatedBy, &createdAt, &updatedAt)
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
		course_id,
		quantity,
		is_deleted,
		created_by,
		updated_by,
		created_at,
		updated_at
	FROM
		class
	WHERE 
		is_deleted = $1
	ORDER BY 
		created_at DESC;
	`
	rows, err := db.Query(query, false)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] query error  %v", err)
		return classLst, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		class := class.Class{}
		var createdAt, updatedAt time.Time
		err = rows.Scan(&class.Id, &class.Code, &class.IdCourse, &class.Quantity, &class.IsDelete, &class.CreatedBy, &class.UpdatedBy, &createdAt, &updatedAt)
		class.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		class.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		classLst = append(classLst, class)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(classLst) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllClass] No Data  %v", err)
		return classLst, errors.New("Không có dữ liệu từ hệ thống")
	}
	return classLst, nil
}
