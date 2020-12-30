package user

import (
	"api-trainning-center/database/usermodel"
	"api-trainning-center/models"
	"database/sql"
	"errors"
)

type IUserService interface {
	CreateAccount(req models.AccountRequest) (models.AccountReponse, error)
	Login(req models.AccountRequest) (models.LoginReponse, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st Store) CreateAccount(req models.AccountRequest) (models.AccountReponse, error) {
	response := models.AccountReponse{}
	// HashPassword hashes password from user input
	hashPassword, err := models.HashPassword(req.PassWord)
	if err != nil {
		return response, err
	}
	req.PassWord = string(hashPassword)

	if err := usermodel.CreateUserByRequest(req, st.db); err != nil {
		return response, errors.New("Username already exists")
	}
	response.Status = true
	return response, nil
}

func (st Store) Login(req models.AccountRequest) (models.LoginReponse, error) {
	resp := models.LoginReponse{}

	if err := usermodel.CheckUserLogin(req, st.db); err != nil {
		return resp, errors.New("Account dose not exists")
	}

	return resp, nil
}
