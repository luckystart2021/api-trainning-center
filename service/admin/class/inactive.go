package class

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreClass) ActiveClass(idClass int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	class, err := models.FindClass(ctx, tx, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	class.IsDeleted = false
	class.UpdatedBy = userName
	_, err = class.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	teacher, err := models.FindTeacher(ctx, tx, class.TeacherID.Int)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] Find Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	if teacher.Status == true {
		logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] Find Teacher Status error : ", err)
		return resp, errors.New("Giáo viên đã tồn tại ở lớp khác, vui lòng chọn giáo viên khác")
	} else {
		teacher.Status = true
		_, err = teacher.Update(ctx, tx, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
			return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	vehicle, err := models.FindVehicle(ctx, tx, class.VehicleID.Int)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindVehicle] Find Vehicle error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if vehicle.Status == true {
		logrus.WithFields(logrus.Fields{}).Error("[FindVehicle] Find Vehicle Status error : ", err)
		return resp, errors.New("Xe đã tồn tại ở lớp khác, vui lòng chọn xe khác")
	} else {
		vehicle.Status = true
		_, err = vehicle.Update(ctx, tx, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[UpdateVehicle] Update Vehicle error : ", err)
			return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	resp.Status = true
	resp.Message = "Cập nhật trạng thái lớp thành công"
	return resp, nil
}

func (st StoreClass) InActiveClass(idClass int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()
	class, err := models.FindClass(ctx, tx, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	class.IsDeleted = true
	class.UpdatedBy = userName
	_, err = class.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	students, err := models.Students(
		qm.Where("class_id = ?", idClass),
	).All(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindStudents] Update Students error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	_, err = students.DeleteAll(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteAllStudents] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	err = UpdateStatusVehicle(ctx, tx, false, int64(class.VehicleID.Int))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateStatusVehicle] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	err = UpdateStatusTeacher(ctx, tx, false, int64(class.TeacherID.Int))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateStatusTeacher] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	resp.Status = true
	resp.Message = "Cập nhật trạng thái lớp thành công"
	return resp, nil
}
