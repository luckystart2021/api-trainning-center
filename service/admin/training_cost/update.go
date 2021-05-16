package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreCost) UpdateCost(id int, req models.TrainingCost, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	ctx := context.Background()
	cost, err := models.FindTrainingCost(ctx, st.db, int64(id))
	if cost == nil {
		return resp, errors.New("Thông tin không tồn tại, vui lòng thử lại")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindTrainingCost] error : ", err)
		return resp, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	cost.Amount = req.Amount
	cost.Type = req.Type
	cost.Note = req.Note
	cost.UpdatedBy = userName
	// cost.ClassID = req.Amount
	// cost.CourseID = req.Amount
	return resp, nil
}
