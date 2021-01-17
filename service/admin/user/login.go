package user

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/utils"
	"errors"

	"github.com/go-redis/redis"
)

func (st Store) Login(req account.AccountRequest, client *redis.Client) (account.LoginReponse, error) {
	response := account.LoginReponse{}
	user, err := account.RetrieveAccountByUserName(req.UserName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	err = account.CheckPasswordHash(req.PassWord, user.PassWord)
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
	user.PassWord = ""
	response.Success = true
	response.Token = token.AccessToken
	response.InfoUser = user
	return response, nil
}
