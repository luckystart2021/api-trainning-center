package fee

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type IFeeService interface {
	UpdateFee(userName string, req models.Fee) (response.MessageResponse, error)
	ShowFees() (models.Fee, error)
}

type StoreFee struct {
	db *sql.DB
}

func NewStoreFee(db *sql.DB) *StoreFee {
	return &StoreFee{
		db: db,
	}
}
