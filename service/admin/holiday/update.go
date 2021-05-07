package holiday

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreHoliday) UpdateHoliday(idHoliday int, req models.Holiday) (response.MessageResponse, error) {
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

	holiday.Name = req.Name
	holiday.Date = req.Date

	rowsAff, err := holiday.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateHoliday] Update Holiday error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if rowsAff > 0 {
		resp.Status = true
		resp.Message = "Cập nhật thông tin ngày nghỉ thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật thông tin ngày nghỉ không thành công"
	}
	return resp, nil
}
