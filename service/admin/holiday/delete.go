package holiday

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreHoliday) DeleteHoliday(idHoliday int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	holiday, err := models.FindHoliday(ctx, st.db, idHoliday)
	if holiday == nil {
		return resp, errors.New("Thông tin ngày lễ không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindHoliday] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	rowsAff, err := holiday.Delete(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteHoliday] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
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
