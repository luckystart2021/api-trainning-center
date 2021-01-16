package user

import (
	"api-trainning-center/models/admin/account"
	"api-trainning-center/validate"
	"errors"
	"log"
)

func (st Store) ResetPassword(userName string) (account.MessageResponse, error) {
	response := account.MessageResponse{}
	user, err := account.RetrieveAccountByUserName(userName, st.db)
	if err != nil {
		return response, err
	}
	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 {
		return response, errors.New("Tên đăng nhập không tồn tại")
	}

	if user.Role == validate.TEACHER {
		hashPassword, err := account.HashPassword("Teacher123@@")
		if err != nil {
			log.Println("hashPassword reset error ", err)
			return response, err
		}
		newPassWord := string(hashPassword)
		if err := account.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
			log.Println("UpdateAccountByRequest reset error ", err)
			return response, err
		}
	}

	if user.Role == validate.EDITOR {
		hashPassword, err := account.HashPassword("Editor123@@")
		if err != nil {
			log.Println("hashPassword reset error ", err)
			return response, err
		}
		newPassWord := string(hashPassword)
		if err := account.UpdateAccountByRequest(user.UserName, newPassWord, st.db); err != nil {
			log.Println("UpdateAccountByRequest reset error ", err)
			return response, err
		}
	}
	response.Status = true
	response.Message = "Reset mật khẩu thành công"
	return response, nil
}
