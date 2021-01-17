package user

import (
	"api-trainning-center/models/admin/account"
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
)

type IUserService interface {
	CreateAccount(req account.AccountRequest) (account.MessageResponse, error)
	Login(req account.AccountRequest, client *redis.Client) (account.LoginReponse, error)
	ChangePassword(req account.ChangeAccountRequest) (account.MessageResponse, error)
	ResetPassword(email string) (account.MessageResponse, error)
	ShowAllAccount() ([]account.User, error)
	ShowAccount(username string) (account.User, error)
	DisableAccountByUserName(username string) (account.MessageResponse, error)
	EnableAccountByUserName(username string) (account.MessageResponse, error)
	UpdateAccountByRequest(req account.AccountRequest) (account.MessageResponse, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st Store) CreateAccount(req account.AccountRequest) (account.MessageResponse, error) {
	response := account.MessageResponse{}

	// HashPassword hashes password from user input
	hashPassword, err := account.HashPassword(req.PassWord)
	if err != nil {
		return response, err
	}
	req.PassWord = string(hashPassword)

	if err := account.CreateUserByRequest(req, st.db); err != nil {
		return response, errors.New("Tên đăng nhập hoặc email hoặc số điện thoại đã tồn tại")
	}
	response.Status = true
	response.Message = "Tạo tài khoản mới thành công"
	return response, nil
}
