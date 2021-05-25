package course

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

func (st StoreCourse) ActiveCourse(idCourse int, userName string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	ctx := context.Background()
	course, err := models.FindCourse(ctx, st.db, idCourse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[ActiveCourse] Update Active Course DB err  %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	course.Status = true
	course.UpdatedBy = userName
	rowsAff, err := course.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ActiveCourse] Update Active Course error : ", err)
		return response, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		response.Status = true
		response.Message = "Đã kích hoạt khóa học thành công"
	} else {
		response.Status = false
		response.Message = "Kích hoạt khóa học không thành công"
	}
	return response, nil
}

func (st StoreCourse) InActiveCourse(idCourse int, userName string) (response.MessageResponse, error) {
	response := response.MessageResponse{}
	ctx := context.Background()

	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	course, err := models.FindCourse(ctx, tx, idCourse)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[InActiveCourse] Update InActive Course DB err  %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	course.Status = false
	course.UpdatedBy = userName
	rowsAff, err := course.Update(ctx, tx, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[InActiveCourse] Update InActive Course error : ", err)
		return response, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	classes, err := models.Classes(
		qm.Where("course_id = ?", idCourse),
		qm.And("teacher_id > ?", 0),
		qm.And("vehicle > ?", 0),
	).All(ctx, tx)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindClass] Find Class DB err  %v", err)
		return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
	}

	for _, class := range classes {
		vehicle, err := models.FindVehicle(ctx, tx, class.VehicleID.Int)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] Find Vehicle DB err  %v", err)
			return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
		vehicle.Status = false
		_, err = vehicle.Update(ctx, tx, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[Update Vehicle] Update Vehicle status error : ", err)
			return response, errors.New("Lỗi hệ thống vui lòng thử lại")
		}

		teachers, err := models.FindTeacher(ctx, tx, class.TeacherID.Int)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] Find Teacher DB err  %v", err)
			return response, errors.New("Lỗi hệ thống, vui lòng thử lại")
		}
		teachers.Status = false
		_, err = teachers.Update(ctx, tx, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[Update Teacher] Update Teacher status error : ", err)
			return response, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAff > 0 {
		response.Status = true
		response.Message = "Hủy kích hoạt khóa học thành công"
	} else {
		response.Status = false
		response.Message = "Hủy kích hoạt khóa học không thành công"
	}
	return response, nil
}
