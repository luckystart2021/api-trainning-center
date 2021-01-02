package account

import (
	"api-trainning-center/database/usermodel"
	"api-trainning-center/models"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"net/http"
	"strings"
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
	response := models.LoginReponse{}
	user, err := usermodel.CheckUserLogin(req, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" {
		return response, errors.New("Login failed, please try again")
	}

	err = models.CheckPasswordHash(req.PassWord, user.PassWord)
	if err != nil {
		return response, errors.New("Login failed, please try again")
	}

	token, err := utils.EncodeAuthToken(user.UserName, user.Role)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Token = token
	return response, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
