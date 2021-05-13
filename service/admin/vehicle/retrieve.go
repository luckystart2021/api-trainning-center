package vehicle

import (
	"api-trainning-center/internal/models"
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreVehicle) ShowVehicles() (models.VehicleSlice, error) {
	vehicles, err := models.Vehicles().All(context.Background(), st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[FindAllVehicles] error : ", err)
		return nil, err
	}
	return vehicles, nil
}

func (st StoreVehicle) ShowVehiclesAvailable() (models.VehicleSlice, error) {
	ctx := context.Background()
	vehicles, err := models.Vehicles(
		qm.Where("status = ?", false),
		qm.And("is_deleted = ?", false),
	).All(ctx, st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowVehiclesAvailable] error : ", err)
		return nil, err
	}
	if vehicles == nil {
		return nil, errors.New("Không có dữ liệu từ hệ thống")
	}
	return vehicles, nil
}

func (st StoreVehicle) ShowVehicle(id int) (models.Vehicle, error) {
	vehicle, err := models.FindVehicle(context.Background(), st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findOneVehicle] error : ", err)
		return models.Vehicle{}, err
	}
	return *vehicle, nil
}
