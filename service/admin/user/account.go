package user

import (
	"api-trainning-center/models/admin"
	"api-trainning-center/validate"
	"database/sql"
	"errors"
	"log"

	"github.com/go-redis/redis"
)

type IUserService interface {
	CreateAccount(req admin.AccountRequest) (admin.Reponse, error)
	Login(req admin.AccountRequest, client *redis.Client) (admin.LoginReponse, error)
	ChangePassword(req admin.ChangeAccountRequest) (admin.Reponse, error)
	ResetPassword(email string) (admin.MessageResponse, error)
	ShowAllAccount() ([]admin.User, error)
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

func (st Store) ChangePassword(req admin.ChangeAccountRequest) (admin.Reponse, error) {
	response := admin.Reponse{}
	user, err := admin.RetrieveAccountByUserName(req.UserName, st.db)
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

	newPassWord := string(hashPassword)

	if err := admin.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
		return response, err
	}

	response.Status = true
	return response, nil
}

func (st Store) ResetPassword(userName string) (admin.MessageResponse, error) {
	response := admin.MessageResponse{}
	user, err := admin.RetrieveAccountByUserName(userName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	if user.Role == validate.TEACHER {
		hashPassword, err := admin.HashPassword("Teacher123@@")
		if err != nil {
			log.Println("hashPassword reset error ", err)
			return response, err
		}
		newPassWord := string(hashPassword)
		if err := admin.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
			log.Println("UpdateAccountByRequest reset error ", err)
			return response, err
		}
	}

	if user.Role == validate.EDITOR {
		hashPassword, err := admin.HashPassword("Editor123@@")
		if err != nil {
			log.Println("hashPassword reset error ", err)
			return response, err
		}
		newPassWord := string(hashPassword)
		if err := admin.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
			log.Println("UpdateAccountByRequest reset error ", err)
			return response, err
		}
	}
	response.Status = true
	response.Message = "Reset mật khẩu thành công"
	return response, nil
}
