package seo

import (
	modelSeo "api-trainning-center/models/admin/seo"
	"api-trainning-center/service/response"
	"database/sql"
)

type ISeoService interface {
	ShowSeoTags() ([]SeoTagsResponse, error)
	ShowDetailSeoTags(id int) (SeoTagsResponse, error)
	ShowSeos() (SeoResponse, error)
	UpdateSeo(id int, req modelSeo.SeoRequest) (response.MessageResponse, error)
	CreateSeoTag(name string) (response.MessageResponse, error)
	UpdateSeoTags(id int, name string) (response.MessageResponse, error)
	DeleteSeoTags(id int) (response.MessageResponse, error)
}

type StoreSeo struct {
	db *sql.DB
}

func NewStoreSeo(db *sql.DB) *StoreSeo {
	return &StoreSeo{
		db: db,
	}
}
