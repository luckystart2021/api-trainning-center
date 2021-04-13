package vehicle

import (
	"api-trainning-center/models/admin/vehicle"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

func (st StoreVehicle) UpdateVehicle(id int, req vehicle.VehicleUpdateRequest, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	count, err := updateVehicleByRequest(st.db, id, req, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateVehicleByRequest] Update Vehicle DB err  %v", err)
		return resp, err
	}
	if count > 0 {
		resp.Status = true
		resp.Message = "Cập nhật xe thành công"
	} else {
		resp.Status = false
		resp.Message = "Cập nhật xe không thành công"
	}
	return resp, nil
}

func updateVehicleByRequest(db *sql.DB, id int, req vehicle.VehicleUpdateRequest, userName string) (int64, error) {
	timeUpdate := time.Now()
	query := `
	UPDATE
		vehicle
	SET
		biensoxe = $1,
		loaixe = $2,
		is_deleted = $3,
		updated_at = $4,
		updated_by = $5,
		is_contract = $7
	WHERE
		id = $6;
	`
	res, err := db.Exec(query, req.BienSoXe, req.LoaiXe, req.IsDeleted, timeUpdate, userName, id, req.IsContract)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[updateVehicleByRequest] update vehicle in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RowsAffected] update vehicle in DB err  %v", err)
		return 0, errors.New("Lỗi hệ thống vui lòng thử lại")
	}

	return rowsAffected, nil
}
