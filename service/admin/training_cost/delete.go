package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreCost) DeleteCost(costID int) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()

	cost, err := models.FindTrainingCost(ctx, st.db, int64(costID))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTrainingCost] FindTrainingCost error : ", err)
		return resp, errors.New("Không có dữ liệu từ hệ thống")
	}

	rowsAff, err := cost.Delete(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DeleteCost] error : ", err)
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
