package user

import (
	"api-trainning-center/models/admin/account"
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
)

type IUserService interface {
	CreateAccount(req account.AccountRequest) (account.Reponse, error)
	Login(req account.AccountRequest, client *redis.Client) (account.LoginReponse, error)
	ChangePassword(req account.ChangeAccountRequest) (account.Reponse, error)
	ResetPassword(email string) (account.MessageResponse, error)
	ShowAllAccount() ([]account.User, error)
	ShowAccount(username string) (account.User, error)
	DeleteAccountByUserName(username string) (account.MessageResponse, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st Store) CreateAccount(req account.AccountRequest) (account.Reponse, error) {
	response := account.Reponse{}

	// HashPassword hashes password from user input
	hashPassword, err := account.HashPassword(req.PassWord)
	if err != nil {
		return response, err
	}
	req.PassWord = string(hashPassword)

	if err := account.CreateUserByRequest(req, st.db); err != nil {
		return response, errors.New("Tên đăng nhập hoặc email đã tồn tại")
	}
	response.Status = true
	return response, nil
}
