package vehicle

import (
	"api-trainning-center/models/admin/vehicle"
	"api-trainning-center/service/response"
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st StoreVehicle) CreateVehicle(req vehicle.VehicleRequest, userName string) (response.MessageResponse, error) {
	resp := response.MessageResponse{}
	if err := createVehicleByRequest(st.db, req, userName); err != nil {
		return resp, err
	}
	resp.Status = true
	resp.Message = "Thêm xe thành công"
	return resp, nil
}

func createVehicleByRequest(db *sql.DB, req vehicle.VehicleRequest, userName string) error {
	query := `
	insert
		into
			vehicle (biensoxe, loaixe, created_by, updated_by, is_contract)
	values($1, $2, $3, $4, $5);
	`
	_, err := db.Exec(query, req.BienSoXe, req.LoaiXe, userName, userName, req.IsContract)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[createVehicleByRequest] Insert vehicle DB err  %v", err)
		return errors.New("Lỗi hệ thống, vui lòng thử lại")
	}
	return nil
}
