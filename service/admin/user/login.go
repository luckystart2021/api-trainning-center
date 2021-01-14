package user

import (
	"api-trainning-center/models/admin"
	"api-trainning-center/utils"
	"errors"

	"github.com/go-redis/redis"
)

func (st Store) Login(req admin.AccountRequest, client *redis.Client) (admin.LoginReponse, error) {
	response := admin.LoginReponse{}
	user, err := admin.RetrieveAccountByUserName(req.UserName, st.db)
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
	user.PassWord = ""
	response.Success = true
	response.Token = token.AccessToken
	response.InfoUser = user
	return response, nil
}
