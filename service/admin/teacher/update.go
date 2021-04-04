package teacher

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreTeacher) UpdateTeacher(id int, req models.Teacher, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()
	teacher, err := models.FindTeacher(ctx, st.db, id)
	if teacher == nil {
		return resp, errors.New("Thông tin giáo viên không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTeacher] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	teacher.Fullname = req.Fullname
	teacher.Sex = req.Sex
	teacher.Dateofbirth = req.Dateofbirth
	teacher.Phone = req.Phone
	teacher.Address = req.Address
	teacher.CMND = req.CMND
	teacher.CNSK = req.CNSK
	teacher.GPLX = req.GPLX
	teacher.ExperienceDriver = req.ExperienceDriver
	teacher.KMSafe = req.KMSafe
	teacher.IsDeleted = req.IsDeleted

	rowsAff, err := teacher.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateTeacher] Update Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin giáo viên thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin giáo viên không thành công"
	}
	return resp, nil
}
