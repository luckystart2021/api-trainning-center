package user

import (
	"api-trainning-center/models/admin/account"
	"errors"
)

func (st Store) ChangePassword(req account.ChangeAccountRequest) (account.Reponse, error) {
	response := account.Reponse{}
	user, err := account.RetrieveAccountByUserName(req.UserName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	err = account.CheckPasswordHash(req.OldPassWord, user.PassWord)
	if err != nil {
		return response, errors.New("Mật khẩu cũ không đúng")
	}

	hashPassword, err := account.HashPassword(req.NewPassWord)
	if err != nil {
		return response, err
	}

	newPassWord := string(hashPassword)

	if err := account.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
		return response, err
	}

	response.Status = true
	return response, nil
}
