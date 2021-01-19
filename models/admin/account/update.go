package account

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// UpdateAccountByRequest executes subscribe to updates from an email address
func UpdateAccountByRequest(userName, newPassWord string, db *sql.DB) error {
	query := `
	UPDATE "users" SET
		password=$1
	WHERE 
		username = $2;`
	_, err := db.Exec(query, newPassWord, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateAccountByRequest] Update DB err  %v", err)
		return err
	}
	return nil
}

func UpdateAccount(req AccountRequest, db *sql.DB) error {
	user := req
	query := `
	UPDATE "users" SET
		email=$2, 
		role=$3, 
		sex =$4, 
		dateofbirth=$5, 
		phone=$6, 
		fullname=$7,
		address=&8
	WHERE 
		username = $1;`
	_, err := db.Exec(query, &user.UserName, &user.Email, &user.Role, &user.Sex, &user.DateOfBirth, &user.Phone, &user.FullName, &user.Address)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[UpdateAccount] Update DB err  %v", err)
		return err
	}
	return nil
}
