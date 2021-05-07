package holiday

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreHoliday) CreateHoliday(req models.Holiday) (response.MessageResponse, error) {
	ctx := context.Background()
	resp := response.MessageResponse{}

	count, err := models.Holidays(
		models.HolidayWhere.Date.EQ(req.Date),
	).Count(ctx, st.db)
	if count > 0 {
		return resp, errors.New("Ngày nghỉ đã tồn tại")
	}

	err = req.Insert(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[CreateHoliday] Create Holiday error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	resp.Status = true
	resp.Message = "Thêm ngày nghỉ thành công"
	return resp, nil
}
