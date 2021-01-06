package account

import (
	"api-trainning-center/models"
	"api-trainning-center/utils"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
)

type IUserService interface {
	CreateAccount(req models.AccountRequest) (models.AccountReponse, error)
	Login(req models.AccountRequest, client *redis.Client) (models.LoginReponse, error)
	ChangeAccount(req models.AccountRequest) (models.AccountReponse, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (st Store) ChangeAccount(req models.AccountRequest) (models.AccountReponse, error) {
	response := models.AccountReponse{}
	return response, nil
}

func (st Store) CreateAccount(req models.AccountRequest) (models.AccountReponse, error) {
	response := models.AccountReponse{}

	// HashPassword hashes password from user input
	hashPassword, err := models.HashPassword(req.PassWord)
	if err != nil {
		return response, err
	}
	req.PassWord = string(hashPassword)

	if err := models.CreateUserByRequest(req, st.db); err != nil {
		return response, errors.New("Username already exists")
	}
	response.Status = true
	return response, nil
}

func (st Store) Login(req models.AccountRequest, client *redis.Client) (models.LoginReponse, error) {
	response := models.LoginReponse{}
	user, err := models.CheckUserLogin(req, st.db)
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
	saveErr := utils.CreateAuth(user.UserName, token, client)
	if saveErr != nil {
		return response, saveErr
	}
	response.Success = true
	response.Token = token.AccessToken
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
