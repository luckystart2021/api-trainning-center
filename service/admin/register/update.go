package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreRegister) UpdateRegister(id int, req models.RegisterGround) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	countDuplicate, err := models.RegisterGrounds(
		qm.Where("start_date = ?", req.StartDate),
		qm.And("end_date = ?", req.EndDate),
		qm.And("ground_number = ?", req.GroundNumber),
		qm.And("class_id = ?", req.ClassID),
	).Count(ctx, st.db)
	if countDuplicate > 0 {
		logrus.WithFields(logrus.Fields{}).Error("[CreateTeacher] Create Teacher error : ", err)
		return resp, errors.New("Thời gian đăng ký đã bị trùng")
	}

	countRegister, err := models.RegisterGrounds(
		qm.Where("end_date > ?", req.StartDate),
	).Count(ctx, st.db)
	if countRegister > 0 {
		logrus.WithFields(logrus.Fields{}).Error("[Count RegisterGrounds] count RegisterGrounds error : ", err)
		return resp, errors.New("Thời gian đăng ký đã có lớp học khác đăng ký")
	}

	registerDB, err := models.FindRegisterGround(ctx, st.db, id)
	if registerDB == nil {
		return resp, errors.New("Thông tin không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindRegisterGround] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	registerDB.StartDate = req.StartDate
	registerDB.EndDate = req.EndDate
	registerDB.ClassID = req.ClassID
	registerDB.TeacherID = req.TeacherID
	registerDB.GroundNumber = req.GroundNumber

	rowsAff, err := registerDB.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Update register] Update register error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin không thành công"
	}
	return resp, nil
}
