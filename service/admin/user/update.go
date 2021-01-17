package user

import (
	"api-trainning-center/models/admin/account"
	"errors"
)

func (st Store) ChangePassword(req account.ChangeAccountRequest) (account.MessageResponse, error) {
	response := account.MessageResponse{}
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
	response.Message = "Cập nhật mật khẩu thành công"
	return response, nil
}

func (st Store) UpdateAccountByRequest(req account.AccountRequest) (account.MessageResponse, error) {
	response := account.MessageResponse{}
	user, err := account.RetrieveAccountByUserName(req.UserName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}
	if err := account.UpdateAccount(req, st.db); err != nil {
		return response, err
	}

	response.Status = true
	response.Message = "Cập nhật thông tin tài khoản thành công"
	return response, nil
}
