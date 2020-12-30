package user

import (
	"api-trainning-center/database"
	"api-trainning-center/models"
	"errors"
)

type IUserService interface {
	CreateAccount(req models.AccountRequest) (models.AccountReponse, error)
}

type Store struct {
	Db database.Database
}

func (st Store) CreateAccount(req models.AccountRequest) (models.AccountReponse, error) {
	response := models.AccountReponse{}
	if err := st.Db.CreateUserByRequest(req); err != nil {
		return response, errors.New("Username already exists")
	}
	response.Status = true
	return response, nil
}
