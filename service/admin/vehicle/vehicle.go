package vehicle

import (
	"api-trainning-center/models/admin/vehicle"
	"api-trainning-center/service/response"
	"database/sql"
)

type IVehicleService interface {
	// ShowStudents() ([]student.Student, error)
	// ShowStudent(idStudent int) (student.Student, error)
	CreateVehicle(req vehicle.VehicleRequest, userName string) (response.MessageResponse, error)
	UpdateVehicle(id int, req vehicle.VehicleUpdateRequest, userName string) (response.MessageResponse, error)
	ShowVehicles() ([]vehicle.Vehicle, error)
	ShowVehicle(id int) (vehicle.FindOneVehicle, error)
}

type StoreVehicle struct {
	db *sql.DB
}

func NewStoreVehicle(db *sql.DB) *StoreVehicle {
	return &StoreVehicle{
		db: db,
	}
}
