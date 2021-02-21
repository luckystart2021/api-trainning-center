package information

import (
	"api-trainning-center/service/response"
	"database/sql"
)

type Information struct {
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Maps        string `json:"maps"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type InformationAdmin struct {
	Id          int    `json:"id"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Maps        string `json:"maps"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Img         string `json:"img"`
	CreatedAt   string `json:"created_at"`
}

type IInformationService interface {
	ShowInformation() (Information, error)
	ShowInformationAdmin() (InformationAdmin, error)
	CreateInformation(address, phone, email, maps, title, description, img string) (response.MessageResponse, error)
	UpdateInformation(idInformationI int, address, phone, email, maps, title, description, img string) (response.MessageResponse, error)
}

type StoreInformation struct {
	db *sql.DB
}

func NewStoreInformation(db *sql.DB) *StoreInformation {
	return &StoreInformation{
		db: db,
	}
}
