package class

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreClass) CreateClass(userName string, idCource, quantity, teacherId, vehicleId int64) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateClass] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	lookupTeacher := make(map[int]int)
	lookupVehicle := make(map[int]int)

	classes, err := models.Classes(
		qm.Where("course_id = ?", idCource),
	).All(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindAllCourses] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	for _, data := range classes {
		lookupTeacher[data.TeacherID.Int] = data.TeacherID.Int
		lookupVehicle[data.VehicleID.Int] = data.VehicleID.Int
	}

	_, ok := lookupTeacher[int(teacherId)]
	if ok {
		return resp, errors.New("Giáo viên đã tồn tại ở lớp khác, vui lòng chọn giáo viên khác")
	}
	_, ok = lookupVehicle[int(vehicleId)]
	if ok {
		return resp, errors.New("Xe đã tồn tại ở lớp khác, vui lòng chọn xe khác")
	}

	if err := CreateClassByRequest(ctx, tx, userName, idCource, quantity, teacherId, vehicleId); err != nil {
		return resp, err
	}

	err = UpdateStatusVehicle(ctx, tx, true, vehicleId)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateStatusVehicle] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	err = UpdateStatusTeacher(ctx, tx, true, teacherId)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateStatusTeacher] err  %v", err)
		return resp, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateAlbum] err  %v", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}
	resp.Status = true
	resp.Message = "Thêm lớp thành công"
	return resp, nil
}

func UpdateStatusTeacher(ctx context.Context, tx *sql.Tx, status bool, teacherId int64) error {
	teacher, err := models.FindTeacher(ctx, tx, int(teacherId))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	teacher.Status = status
	_, err = teacher.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Vehicle error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return nil
}

func UpdateStatusVehicle(ctx context.Context, tx *sql.Tx, status bool, vehicleId int64) error {
	vehicle, err := models.FindVehicle(ctx, tx, int(vehicleId))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] err  %v", err)
		return errors.New("Không có dữ liệu từ hệ thống")
	}
	vehicle.Status = status
	_, err = vehicle.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateVehicle] Update Vehicle error : ", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return nil
}

func CreateClassByRequest(ctx context.Context, tx *sql.Tx, userName string, idCource, quantity, teacherId, vehicleId int64) error {
	query := `
	INSERT INTO class
		(code, course_id, quantity, created_by, updated_by,teacher_id, vehicle_id)
	(
	SELECT
		CONCAT('L-', COUNT(*)+1), $1, $2, $3, $4, $5, $6
	FROM
		"class");
	`
	_, err := tx.ExecContext(ctx, query, idCource, quantity, userName, userName, teacherId, vehicleId)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateClassByRequest]Insert class DB err  %v", err)
		tx.Rollback()
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
