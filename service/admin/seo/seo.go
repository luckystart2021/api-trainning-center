package seo

import "database/sql"

type ISeoService interface {
	ShowSeos() (Seo, error)
}

type StoreSeo struct {
	db *sql.DB
}

func NewStoreSeo(db *sql.DB) *StoreSeo {
	return &StoreSeo{
		db: db,
	}
}
