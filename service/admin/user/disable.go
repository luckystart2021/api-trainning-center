package user

import (
	"api-trainning-center/models/admin/account"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st Store) DisableAccountByUserName(username string) (account.MessageResponse, error) {
	resp := account.MessageResponse{}
	user, err := account.RetrieveAccountByUserName(username, st.db)

	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 || len(username) > 50 {
		return resp, errors.New("Tên đăng nhập không tồn tại hoặc đã bị khóa")
	}

	if err := account.DisableAccountByUserName(username, st.db); err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DisableAccountByUserName] from db error : ", err)
		return resp, err
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[DisableAccountByUserName] error : ", err)
		return resp, err
	}
	resp.Status = true
	resp.Message = "Vô hiệu hóa tài khoản thành công"
	return resp, nil
}
