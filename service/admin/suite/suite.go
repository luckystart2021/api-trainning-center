package suite

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type ISuiteTestService interface {
	GenarateSuiteTest(number, rank string) (response.MessageResponse, error)
}

type StoreSuiteTest struct {
	db *sql.DB
}

func NewStoreSuiteTest(db *sql.DB) *StoreSuiteTest {
	return &StoreSuiteTest{
		db: db,
	}
}
