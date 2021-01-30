package about

import "database/sql"

type About struct {
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type IAboutService interface {
	ShowAbout() ([]About, error)
}

type StoreAbout struct {
	db *sql.DB
}

func NewStoreAbout(db *sql.DB) *StoreAbout {
	return &StoreAbout{
		db: db,
	}
}
