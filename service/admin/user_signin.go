package admin

import (
	"api-trainning-center/database"
	"api-trainning-center/models"
)

type IUserService interface {
	CreateAccount(req *models.AccountRequest) (*models.AccountReponse, error)
}
type Store struct {
	Db database.Database
}

func (st Store) CreateAccount(req *models.AccountRequest) (*models.AccountReponse, error) {
	response := &models.AccountReponse{}
	if err := st.Db.CreateUserByEmail(req.UserName, req.PassWord, req.Email, req.Role); err != nil {
		return response, err
	}
	response.Status = true
	return response, nil
}
