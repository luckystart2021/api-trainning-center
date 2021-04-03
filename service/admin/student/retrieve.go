package student

import (
	"api-trainning-center/models/admin/student"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
)

func (st StoreStudent) ShowStudents() ([]student.Student, error) {
	students, err := FindAllStudents(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowStudents] error : ", err)
		return nil, err
	}
	return students, nil
}

func (st StoreStudent) ShowStudent(idStudent int) (student.Student, error) {
	sdt := student.Student{}
	student, err := FindOneStudent(st.db, idStudent)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowStudent] error : ", err)
		return sdt, err
	}
	return student, nil
}

func FindOneStudent(db *sql.DB, idStudent int) (student.Student, error) {
	student := student.Student{}
	query := `
	SELECT
		id,
		code,
		sex,
		dateofbirth,
		phone,
		address,
		fullname,
		class_id,
		created_by,
		created_at,
		updated_by,
		updated_at,
		cmnd,
		cnsk,
		gplx,
		experience_driver,
		km_safe
	FROM
		student
	WHERE
		id = $1;
	`
	rows := db.QueryRow(query, idStudent)
	var createdAt, updatedAt time.Time
	var gplx null.String
	err := rows.Scan(&student.Id, &student.Code, &student.Sex, &student.DateOfBirth, &student.Phone,
		&student.Address, &student.FullName, &student.IdClass, &student.CreatedBy,
		&createdAt, &student.UpdatedBy, &updatedAt, &student.CMND,
		&student.CNSK, &gplx, &student.Exp, &student.NumberOfKm)
	student.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	student.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	if gplx.Valid {
		student.GPLX = gplx.String
	}

	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneStudent] No Data  %v", err)
		return student, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneStudent] Scan error  %v", err)
	}
	return student, nil
}

func FindAllStudents(db *sql.DB) ([]student.Student, error) {
	students := []student.Student{}
	query := `
	SELECT
		id,
		code,
		sex,
		dateofbirth,
		phone,
		address,
		fullname,
		class_id,
		created_by,
		created_at,
		updated_by,
		updated_at,
		cmnd,
		cnsk,
		gplx,
		experience_driver,
		km_safe
	FROM
		student;`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllStudents] query error  %v", err)
		return students, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		student := student.Student{}
		var createdAt, updatedAt time.Time
		var gplx null.String
		err = rows.Scan(&student.Id, &student.Code, &student.Sex, &student.DateOfBirth, &student.Phone,
			&student.Address, &student.FullName, &student.IdClass, &student.CreatedBy, &createdAt,
			&student.UpdatedBy, &updatedAt, &student.CMND,
			&student.CNSK, &gplx, &student.Exp, &student.NumberOfKm)
		student.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		student.UpdatedAt = utils.TimeIn(updatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		if gplx.Valid {
			student.GPLX = gplx.String
		}
		students = append(students, student)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneStudent] Scan error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllStudents] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(students) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllStudents] No Data  %v", err)
		return students, errors.New("Không có dữ liệu từ hệ thống")
	}
	return students, nil
}
