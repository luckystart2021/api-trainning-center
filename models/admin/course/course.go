package course

import "time"

type Course struct {
	Id             int       `json:"id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	GraduationDate time.Time `json:"graduation_date"`
	TestDate       time.Time `json:"test_date"`
	TrainingSystem string    `json:"training_system"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
}
