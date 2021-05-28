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

func (st StoreRegister) CreateRegister(req models.RegisterGround) (response.MessageResponse, error) {
	ctx := context.Background()
	resp := response.MessageResponse{}

	countRegister, err := models.RegisterGrounds(
		qm.Where("end_date > ?", req.StartDate),
	).Count(ctx, st.db)
	if countRegister > 0 {
		logrus.WithFields(logrus.Fields{}).Error("[CreateTeacher] Create Teacher error : ", err)
		return resp, errors.New("Thời gian đăng ký đã có lớp học khác đăng ký")
	}

	err = req.Insert(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CreateTeacher] Create Teacher error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	resp.Status = true
	resp.Message = "Đăng ký thành công"
	return resp, nil
}
