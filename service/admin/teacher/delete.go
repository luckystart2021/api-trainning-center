package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreTeacher) InActive(idTeacher int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	teacher, err := models.FindTeacher(ctx, st.db, idTeacher)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] Find Teacher DB err  %v", err)
		return resp, err
	}
	var rowsAff int64
	if teacher.Status {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] Find Teacher DB err  %v", err)
		return resp, errors.New("Giáo viên đang tồn tại trong lớp học, vui lòng cập nhật giáo viên khác trước khi xóa")
	} else {
		teacher.IsDeleted = true
		teacher.UpdatedBy = userName
		rowsAff, err = teacher.Update(ctx, st.db, boil.Infer())
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

func (st StoreTeacher) Active(idTeacher int, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	teacher, err := models.FindTeacher(ctx, st.db, idTeacher)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindTeacher] Find Teacher DB err  %v", err)
		return resp, err
	}
	teacher.IsDeleted = false
	teacher.UpdatedBy = userName
	rowsAff, err := teacher.Update(ctx, st.db, boil.Infer())
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
