package user

import (
	"api-trainning-center/models/admin"
	"api-trainning-center/utils"
	"database/sql"
	"errors"

	"github.com/go-redis/redis"
)

type IUserService interface {
	CreateAccount(req admin.AccountRequest) (admin.Reponse, error)
	Login(req admin.AccountRequest, client *redis.Client) (admin.LoginReponse, error)
	ChangePassword(req admin.ChangeAccountRequest) (admin.Reponse, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st Store) CreateAccount(req admin.AccountRequest) (admin.Reponse, error) {
	response := admin.Reponse{}

	// HashPassword hashes password from user input
	hashPassword, err := admin.HashPassword(req.PassWord)
	if err != nil {
		return response, err
	}
	req.PassWord = string(hashPassword)

	if err := admin.CreateUserByRequest(req, st.db); err != nil {
		return response, errors.New("Tên đăng nhập hoặc email đã tồn tại")
	}
	response.Status = true
	return response, nil
}

func (st Store) Login(req admin.AccountRequest, client *redis.Client) (admin.LoginReponse, error) {
	response := admin.LoginReponse{}
	user, err := admin.CheckUserLogin(req.UserName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	err = admin.CheckPasswordHash(req.PassWord, user.PassWord)
	if err != nil {
		return response, errors.New("Đăng nhập thất bại")
	}

	token, err := utils.EncodeAuthToken(user.UserName, user.Role)
	if err != nil {
		return response, err
	}

	saveErr := utils.CreateAuth(token, client)
	if saveErr != nil {
		return response, saveErr
	}

	response.Success = true
	response.Token = token.AccessToken
	response.UserID = user.UserName
	return response, nil
}

func (st Store) ChangePassword(req admin.ChangeAccountRequest) (admin.Reponse, error) {
	response := admin.Reponse{}
	user, err := admin.CheckUserLogin(req.UserName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	err = admin.CheckPasswordHash(req.OldPassWord, user.PassWord)
	if err != nil {
		return response, errors.New("Mật khẩu cũ không đúng")
	}

	hashPassword, err := admin.HashPassword(req.NewPassWord)
	if err != nil {
		return response, err
	}

	req.NewPassWord = string(hashPassword)

	if err := admin.UpdateAccountByRequest(req, st.db); err != nil {
		return response, err
	}

	response.Status = true
	return response, nil
}
