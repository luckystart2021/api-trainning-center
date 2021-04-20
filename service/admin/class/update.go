package class

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreClass) UpdateClass(idClass int, userName string, idCource, quantity int64, isDeleted bool, teacherId, vehicleId int64) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateClass] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	err = CheckAndUpdateStatusTeacher(ctx, tx, true, teacherId, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CheckAndUpdateStatusTeacher] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	err = CheckAndUpdateStatusVehicle(ctx, tx, true, vehicleId, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateStatusVehicle] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	if err := UpdateClassByRequest(ctx, tx, idClass, userName, idCource, quantity, isDeleted, teacherId, vehicleId); err != nil {
		return resp, err
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbum] err  %v", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	resp.Status = true
	resp.Message = "Cập nhật lớp thành công"
	return resp, nil
}

func CheckAndUpdateStatusTeacher(ctx context.Context, tx *sql.Tx, status bool, teacherId int64, idClass int) error {
	class, err := models.FindClass(ctx, tx, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindClass] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	teacher, err := models.FindTeacher(ctx, tx, int(teacherId))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	teacherOld, err := models.FindTeacher(ctx, tx, class.TeacherID.Int)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}

	if int64(class.VehicleID.Int) == int64(teacher.ID) {
		return nil
	}

	teacher.Status = status
	teacherOld.Status = false
	_, err = teacherOld.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Vehicle Old error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	_, err = teacher.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Vehicle error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return nil
}

func CheckAndUpdateStatusVehicle(ctx context.Context, tx *sql.Tx, status bool, vehicleId int64, idClass int) error {
	class, err := models.FindClass(ctx, tx, idClass)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindClass] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	vehicle, err := models.FindVehicle(ctx, tx, int(vehicleId))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	vehicleOld, err := models.FindVehicle(ctx, tx, class.VehicleID.Int)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}

	if int64(class.VehicleID.Int) == int64(vehicle.ID) {
		return nil
	}

	vehicle.Status = status
	vehicleOld.Status = false
	_, err = vehicleOld.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateVehicle] Update Vehicle Old error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	_, err = vehicle.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateVehicle] Update Vehicle error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return nil
}

func UpdateClassByRequest(ctx context.Context, tx *sql.Tx, idClass int, userName string, idCource, quantity int64, isDeleted bool, teacherId, vehicleId int64) error {
	timeUpdate := time.Now()
	query := `
	UPDATE
		class
	SET
		course_id = $1,
		quantity = $2,
		updated_by = $3,
		updated_at = $4,
		is_deleted = $6,
		teacher_id = $7,
		vehicle_id = $8
	WHERE
		id = $5;
	`
	_, err := tx.ExecContext(ctx, query, idCource, quantity, userName, timeUpdate, idClass, isDeleted, teacherId, vehicleId)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateClassByRequest] update class DB err  %v", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
