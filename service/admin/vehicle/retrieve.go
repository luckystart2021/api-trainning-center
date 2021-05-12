package vehicle

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/models/admin/vehicle"
	"api-trainning-center/utils"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (st StoreVehicle) ShowVehicles() ([]vehicle.Vehicle, error) {
	vehicles, err := findAllVehicles(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findAllVehicles] error : ", err)
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

func (st StoreVehicle) ShowVehicle(id int) (vehicle.FindOneVehicle, error) {
	v := vehicle.FindOneVehicle{}
	vehicle, err := findOneVehicle(st.db, id)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[findOneVehicle] error : ", err)
		return v, err
	}
	return vehicle, nil
}

func findOneVehicle(db *sql.DB, id int) (vehicle.FindOneVehicle, error) {
	vehicle := vehicle.FindOneVehicle{}
	query := `
	SELECT
		id,
		biensoxe,
		loaixe,
		status,
		is_deleted,
		created_by,
		created_at,
		updated_at,
		updated_by,
		is_contract
	FROM
		vehicle
	WHERE
		id = $1
	`
	rows := db.QueryRow(query, id)
	var createdAt, updatedAt time.Time
	err := rows.Scan(&vehicle.Id, &vehicle.BienSoXe, &vehicle.LoaiXe, &vehicle.Status, &vehicle.IsDeleted, &vehicle.CreatedBy, &createdAt, &updatedAt, &vehicle.UpdatedBy, &vehicle.IsContract)
	vehicle.CreatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	vehicle.UpdatedAt = utils.TimeIn(createdAt, utils.TIMEZONE, utils.LAYOUTTIMEDDMMYYYYHHMMSS)
	if err == sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{}).Errorf("[findOneVehicle] No Data  %v", err)
		return vehicle, errors.New("Không có dữ liệu từ hệ thống")
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findOneVehicle] Scan error  %v", err)
	}
	return vehicle, nil
}

func findAllVehicles(db *sql.DB) ([]vehicle.Vehicle, error) {
	vehicles := []vehicle.Vehicle{}
	query := `
	SELECT
		id,
		biensoxe,
		loaixe,
		status,
		is_deleted,
		is_contract
	FROM
		vehicle;
	`
	rows, err := db.Query(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAllVehicles] query error  %v", err)
		return vehicles, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	defer rows.Close()
	for rows.Next() {
		vehicle := vehicle.Vehicle{}
		err = rows.Scan(&vehicle.Id, &vehicle.BienSoXe, &vehicle.LoaiXe, &vehicle.Status, &vehicle.IsDeleted, &vehicle.IsContract)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[findAllVehicles] Scan error  %v", err)
			return vehicles, errors.New("Lỗi hệ thống vui lòng thử lại")
		}
		vehicles = append(vehicles, vehicle)
	}
	err = rows.Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAllVehicles] Rows error  %v", err)
		return nil, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	if len(vehicles) == 0 {
		logrus.WithFields(logrus.Fields{}).Errorf("[findAllVehicles] No Data  %v", err)
		return vehicles, errors.New("Không có dữ liệu từ hệ thống")
	}
	return vehicles, nil
}
