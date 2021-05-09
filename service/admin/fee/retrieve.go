package fee

import (
	"api-trainning-center/internal/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreFee) ShowFees() (models.Fee, error) {
	ctx := context.Background()
	feeR := models.Fee{}
	fees, err := models.Fees(
		qm.Limit(1),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindFees] error : ", err)
		return feeR, err
	}
	if fees == nil {
		return feeR, errors.New("Không có dữ liệu từ hệ thống")
	}
	for _, data := range fees {
		feeR = *data
	}

	return feeR, nil
}
