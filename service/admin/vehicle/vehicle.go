package vehicle

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/models/admin/vehicle"
	"api-trainning-center/service/response"
	"database/sql"
)

type IVehicleService interface {
	CreateVehicle(req vehicle.VehicleRequest, userName string) (response.MessageResponse, error)
	UpdateVehicle(id int, req vehicle.VehicleUpdateRequest, userName string) (response.MessageResponse, error)
	ShowVehicles() (models.VehicleSlice, error)
	ShowVehicle(id int) (models.Vehicle, error)
	ShowVehiclesAvailable() (models.VehicleSlice, error)
	InActiveVehicle(id int, userName string) (response.MessageResponse, error)
	ActiveVehicle(id int, userName string) (response.MessageResponse, error)
}

type StoreVehicle struct {
	db *sql.DB
}

func NewStoreVehicle(db *sql.DB) *StoreVehicle {
	return &StoreVehicle{
		db: db,
	}
}
