package seo

import (
	modelSeo "api-trainning-center/models/admin/seo"
	"api-trainning-center/service/response"
	"database/sql"
)

type ISeoService interface {
	ShowSeos() (SeoResponse, error)
	UpdateSeo(id int, req modelSeo.SeoRequest) (response.MessageResponse, error)
}

type StoreSeo struct {
	db *sql.DB
}

func NewStoreSeo(db *sql.DB) *StoreSeo {
	return &StoreSeo{
		db: db,
	}
}
