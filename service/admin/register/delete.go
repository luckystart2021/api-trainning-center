package register

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreRegister) DeleteRegister(id int) (response.MessageResponse, error) {
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

	rowsAff, err := register.Delete(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteStudent] error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Xóa thành công"
	} else {
		resp.Status = false
		resp.Message = "Xóa không thành công"
	}
	return resp, nil
}
