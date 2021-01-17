package user

import (
	"api-trainning-center/models/admin/account"
	"errors"

	"github.com/sirupsen/logrus"
)

func (st Store) ShowAllAccount() ([]account.User, error) {
	users, err := account.RetrieveAccounts(st.db)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowAllAccount] error : ", err)
		return []account.User{}, err
	}

	return users, nil
}

func (st Store) ShowAccount(username string) (account.User, error) {
	user, err := account.RetrieveAccountByUserName(username, st.db)

	// user is not registered
	if user.UserName == "" || len(user.UserName) == 0 || len(user.UserName) > 50 {
		return account.User{}, errors.New("Tên đăng nhập không tồn tại hoặc đã bị khóa")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Error("[ShowAccount] error : ", err)
		return account.User{}, err
	}

	return user, nil
}
