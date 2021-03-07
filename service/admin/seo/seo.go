package seo

import "database/sql"

type ISeoService interface {
	
}

type StoreSeo struct {
	db *sql.DB
}

func NewStoreSeo(db *sql.DB) *StoreSeo {
	return &StoreSeo{
		db: db,
	}
}
