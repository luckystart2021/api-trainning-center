package account

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func DisableAccountByUserName(userName string, db *sql.DB) error {
	query := `
	UPDATE "users" SET
		is_delete=$1
	WHERE 
		username = $2;`
	_, err := db.Exec(query, INACTIVE, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[DeleteAccountByUserName] Delete DB err  %v", err)
		return err
	}
	return nil
}
