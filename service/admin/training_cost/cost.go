package training_cost

import (
	"api-trainning-center/internal/models"
	"api-trainning-center/service/response"
	"database/sql"
)

type ICostService interface {
	CreateCost(req models.TrainingCost, courseID, classID int, userName string) (response.MessageResponse, error)
	UpdateCost(id int, req models.TrainingCost, userName string, courseID, classID int) (response.MessageResponse, error)
	ShowCost(courseID int) ([]TrainingCost, error)
	ShowDetailCost(costID int) (TrainingCost, error)
	DeleteCost(costID int) (response.MessageResponse, error)
	ShowCostByClass(classID int) ([]TrainingCost, error)
}

type StoreCost struct {
	db *sql.DB
}

func NewStoreCost(db *sql.DB) *StoreCost {
	return &StoreCost{
		db: db,
	}
}
