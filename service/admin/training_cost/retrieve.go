package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/utils"
	"context"

	"github.com/leekchan/accounting"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type TrainingCost struct {
	ID        int64  `json:"id"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Note      string `json:"note"`
	ClassID   int    `json:"class_id"`
	CourseID  int    `json:"course_id"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

func (st StoreCost) ShowCost(course int) ([]TrainingCost, error) {
	trainingCosts := []TrainingCost{}
	ctx := context.Background()
	costs, err := models.TrainingCosts(
		qm.Where("course_id = ?", course),
		qm.OrderBy("created_at DESC"),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[TrainingCosts] error : ", err)
		return nil, err
	}
	for _, cost := range costs {
		trainingCost := TrainingCost{}
		trainingCost.ID = cost.ID
		ac := accounting.Accounting{Precision: 0}
		trainingCost.Amount = ac.FormatMoney(cost.Amount)
		trainingCost.Type = cost.Type.String
		trainingCost.Note = cost.Note.String
		trainingCost.ClassID = cost.ClassID
		trainingCost.CourseID = cost.CourseID
		trainingCost.CreatedAt = utils.TimeIn(cost.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		trainingCost.CreatedBy = cost.CreatedBy
		trainingCost.UpdatedAt = utils.TimeIn(cost.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
		trainingCost.UpdatedBy = cost.UpdatedBy

		trainingCosts = append(trainingCosts, trainingCost)
	}
	return trainingCosts, nil
}

func (st StoreCost) ShowDetailCost(costID int) (TrainingCost, error) {
	ctx := context.Background()
	trainingCost := TrainingCost{}

	cost, err := models.FindTrainingCost(ctx, st.db, int64(costID))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[TrainingCosts] error : ", err)
		return trainingCost, err
	}

	trainingCost.ID = cost.ID
	ac := accounting.Accounting{Precision: 0}
	trainingCost.Amount = ac.FormatMoney(cost.Amount)
	trainingCost.Type = cost.Type.String
	trainingCost.Note = cost.Note.String
	trainingCost.ClassID = cost.ClassID
	trainingCost.CourseID = cost.CourseID
	trainingCost.CreatedAt = utils.TimeIn(cost.CreatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	trainingCost.CreatedBy = cost.CreatedBy
	trainingCost.UpdatedAt = utils.TimeIn(cost.UpdatedAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	trainingCost.UpdatedBy = cost.UpdatedBy

	return trainingCost, nil
}
