package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreRegister) UpdateRegister(id int, req models.RegisterGround) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	register, err := models.FindRegisterGround(ctx, st.db, id)
	if register == nil {
		return resp, errors.New("Thông tin không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindRegisterGround] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	register.StartDate = req.StartDate
	register.GroundNumber = req.GroundNumber

	rowsAff, err := register.Update(ctx, st.db, boil.Infer())
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
