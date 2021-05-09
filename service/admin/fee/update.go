package fee

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/boil"
)

func (st StoreFee) UpdateFee(userName string, req models.Fee) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	feeDb, err := models.Fees().One(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindFee] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	fee, err := models.FindFee(ctx, st.db, feeDb.ID)
	if fee == nil {
		return resp, errors.New("Thông tin học phí không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindFee] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	fee.Amount = req.Amount
	fee.UpdatedBy = userName

	rowsAff, err := fee.Update(ctx, st.db, boil.Infer())
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[UpdateFee] Update Fee error : ", err)
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
