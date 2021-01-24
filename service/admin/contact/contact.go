package contact

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type IContactService interface {
	CreateContact(fullName, phone, email, message, subject string) (response.MessageResponse, error)
	ShowContacts() ([]Contact, error)
}

type StoreContact struct {
	db *sql.DB
}

func NewStoreContact(db *sql.DB) *StoreContact {
	return &StoreContact{
		db: db,
	}
}
