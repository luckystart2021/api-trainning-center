package register

import (
	"api-trainning-center/internal/models"
	"context"

	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreRegister) ShowRegisterByClassId(classID int) (models.RegisterGroundSlice, error) {
	ctx := context.Background()
	register, err := models.RegisterGrounds(
		qm.Where("class_id = ?", classID),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find RegisterGrounds] error : ", err)
		return nil, err
	}
	if register == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return register, nil
}

func (st StoreRegister) ShowRegisterById(Id int) (models.RegisterGround, error) {
	ctx := context.Background()
	register, err := models.FindRegisterGround(ctx, st.db, Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[Find RegisterGrounds By Id] error : ", err)
		return models.RegisterGround{}, err
	}
	if register == nil {
		return models.RegisterGround{}, errors.New("Không có dữ liệu từ hệ thống")
	}

	return *register, nil
}
