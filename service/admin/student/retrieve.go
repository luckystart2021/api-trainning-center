package student

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/models/admin/student"
	"api-trainning-center/utils"
	"context"

	"github.com/leekchan/accounting"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreStudent) SearchStudentInformation(codeStudent string) (student.Student, error) {
	ctx := context.Background()
	student := student.Student{}

	amountDB, err := models.Fees().One(context.Background(), st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find Fee] error : ", err)
		return student, err
	}

	data, err := models.Students(
		qm.Where("code = ?", codeStudent),
	).One(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find Students by code] error : ", err)
		return student, err
	}

	student.Id = int64(data.ID)
	student.Code = data.Code
	student.Email = data.Email.String
	student.Sex = data.Sex
	student.DateOfBirth = data.Dateofbirth
	student.Phone = data.Phone
	student.Address = data.Address
	student.FullName = data.Fullname
	student.IdClass = int64(data.ClassID)
	student.CreatedAt = utils.TimeIn(data.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	student.CreatedBy = data.CreatedBy
	student.UpdatedAt = utils.TimeIn(data.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	student.UpdatedBy = data.UpdatedBy
	student.CMND = data.CMND
	student.CNSK = data.CNSK
	student.GPLX = data.GPLX.String
	student.Exp = data.ExperienceDriver
	student.NumberOfKm = data.KMSafe
	ac := accounting.Accounting{Precision: 0}
	student.AmountComplete = ac.FormatMoney(data.Amount.Float64)
	student.AmountRemain = ac.FormatMoney(amountDB.Amount - data.Amount.Float64)
	student.KetQua = "Không đậu"
	if data.KetQua.Bool {
		student.KetQua = "Đậu"
	}
	if !data.KetQua.Valid {
		student.KetQua = "Chưa có kết quả"
	}

	return student, nil
}

func (st StoreStudent) ShowStudents() ([]student.Student, error) {
	students := []student.Student{}
	ctx := context.Background()
	studentsDB, err := models.Students(
		qm.OrderBy("id DESC"),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find Students] error : ", err)
		return nil, err
	}

	amountDB, err := models.Fees().One(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find Fee] error : ", err)
		return nil, err
	}

	for _, data := range studentsDB {
		student := student.Student{}
		student.Id = int64(data.ID)
		student.Email = data.Email.String
		student.Code = data.Code
		student.Sex = data.Sex
		student.DateOfBirth = data.Dateofbirth
		student.Phone = data.Phone
		student.Address = data.Address
		student.FullName = data.Fullname
		student.IdClass = int64(data.ClassID)
		student.CreatedAt = utils.TimeIn(data.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		student.CreatedBy = data.CreatedBy
		student.UpdatedAt = utils.TimeIn(data.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		student.UpdatedBy = data.UpdatedBy
		student.CMND = data.CMND
		student.CNSK = data.CNSK
		student.GPLX = data.GPLX.String
		student.Exp = data.ExperienceDriver
		student.NumberOfKm = data.KMSafe
		ac := accounting.Accounting{Precision: 0}
		student.AmountComplete = ac.FormatMoney(data.Amount.Float64)
		student.AmountRemain = ac.FormatMoney(amountDB.Amount - data.Amount.Float64)
		student.DiemLyThuyet = data.DiemLyThuyet.String
		student.DiemThucHanh = data.DiemThucHanh.String
		student.KetQua = "Không đậu"
		if data.KetQua.Bool {
			student.KetQua = "Đậu"
		}
		if !data.KetQua.Valid {
			student.KetQua = "Chưa có kết quả"
		}

		students = append(students, student)
	}

	return students, nil
}

func (st StoreStudent) ShowStudent(idStudent int) (student.Student, error) {
	student := student.Student{}

	amountDB, err := models.Fees().One(context.Background(), st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find Fee] error : ", err)
		return student, err
	}

	data, err := models.FindStudent(context.Background(), st.db, idStudent)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindStudent] error : ", err)
		return student, err
	}

	student.Id = int64(data.ID)
	student.Code = data.Code
	student.Email = data.Email.String
	student.Sex = data.Sex
	student.DateOfBirth = data.Dateofbirth
	student.Phone = data.Phone
	student.Address = data.Address
	student.FullName = data.Fullname
	student.IdClass = int64(data.ClassID)
	student.CreatedAt = utils.TimeIn(data.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	student.CreatedBy = data.CreatedBy
	student.UpdatedAt = utils.TimeIn(data.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	student.UpdatedBy = data.UpdatedBy
	student.CMND = data.CMND
	student.CNSK = data.CNSK
	student.GPLX = data.GPLX.String
	student.Exp = data.ExperienceDriver
	student.NumberOfKm = data.KMSafe
	ac := accounting.Accounting{Precision: 0}
	student.AmountComplete = ac.FormatMoney(data.Amount.Float64)
	student.AmountRemain = ac.FormatMoney(amountDB.Amount - data.Amount.Float64)
	student.KetQua = "Không đậu"
	if data.KetQua.Bool {
		student.KetQua = "Đậu"
	}
	if !data.KetQua.Valid {
		student.KetQua = "Chưa có kết quả"
	}

	return student, nil
}
