package holiday

import (
	"api-trainning-center/internal/models"
	"context"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreHoliday) ShowHolidays() (models.HolidaySlice, error) {
	ctx := context.Background()
	holidays, err := models.Holidays(
		qm.OrderBy("date DESC"),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findHolidays] error : ", err)
		return nil, err
	}
	if holidays == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return holidays, nil
}

func (st StoreHoliday) ShowHoliday(idHoliday int) (models.Holiday, error) {
	ctx := context.Background()
	holiday, err := models.FindHoliday(ctx, st.db, idHoliday)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[FindOneHoliday] No Data  %v", err)
		return models.Holiday{}, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindOneHoliday] error : ", err)
		return models.Holiday{}, err
	}

	return *holiday, nil
}
