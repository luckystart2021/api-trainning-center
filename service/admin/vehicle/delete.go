package vehicle

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreVehicle) InActiveVehicle(id int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	vehicle, err := models.FindVehicle(ctx, st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] Find Vehicle DB err  %v", err)
		return resp, err
	}
	var rowsAff int64
	if vehicle.Status {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] Find Vehicle DB err  %v", err)
		return resp, errors.New("Xe đang tồn tại trong lớp học, vui lòng cập nhật xe khác trước khi xóa")
	} else {
		vehicle.IsDeleted = true
		vehicle.UpdatedBy = userName
		rowsAff, err = vehicle.Update(ctx, st.db, boil.Infer())
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
			return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
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

func (st StoreVehicle) ActiveVehicle(id int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()
	vehicle, err := models.FindVehicle(ctx, st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindVehicle] Find Vehicle DB err  %v", err)
		return resp, err
	}
	vehicle.IsDeleted = false
	vehicle.UpdatedBy = userName
	rowsAff, err := vehicle.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
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
