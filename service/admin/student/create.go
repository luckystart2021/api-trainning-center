package student

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreStudent) CreateStudent(sex, dayOfBirth, phone, address, fullName, userName string, idClass int, cmnd string, cnsk bool,
	gplx string, exp, numberKm int, amount float64) (response.MessageResponse, error) {
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
	// query := `
	// INSERT INTO student
	// 	(code, sex, dateofbirth, phone, address, fullname, class_id, created_by, updated_by, cmnd, cnsk, gplx, experience_driver, km_safe, amount)
	// (
	// SELECT
	// 	CONCAT('HV', COUNT(*)+1), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
	// FROM
	// 	student
	// );
	// `
	// _, err := db.Exec(query, sex, dayOfBirth, phone, address, fullName, idClass, userName, userName, cmnd, cnsk, gplx, exp, numberKm, amount)
	// if err != nil {
	// 	logrus.WithFields(logrus.Fields{}).Errorf("[CreateStudentByRequest]Insert student DB err  %v", err)
	// 	return errors.New("Lỗi hệ thống, vui lòng thử lại")
	// }
	// return nil
	ctx := context.Background()

	countSt, err := models.Students().Count(ctx, db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CountStudent] Count Student error : ", err)
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	student := models.Student{}
	student.Code = "HV-" + strconv.Itoa(int(countSt)+1)
	student.ClassID = idClass
	student.Sex = sex
	student.Dateofbirth = dayOfBirth
	student.Phone = phone
	student.Address = address
	student.Fullname = fullName
	student.CMND = cmnd
	student.CNSK = cnsk
	student.GPLX = null.StringFrom(gplx)
	student.ExperienceDriver = exp
	student.KMSafe = numberKm
	student.Amount = null.Float64From(amount)
	// student.DiemLyThuyet = null.StringFrom(diemLythuyet)
	// student.DiemThucHanh = null.StringFrom(diemThucHanh)
	// student.KetQua = null.BoolFrom(ketQua)
	student.CreatedBy = userName
	student.UpdatedBy = userName

	err = student.Insert(ctx, db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CreateStudent] Create Student error : ", err)
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return nil

}
