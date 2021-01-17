package account

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

func EnableAccountByUserName(userName string, db *sql.DB) error {
	query := `
	UPDATE "user" SET
		is_delete=$1
	WHERE 
		username = $2;`
	_, err := db.Exec(query, ACTIVE, userName)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[EnableAccountByUserName] Update is delete DB err  %v", err)
		return err
	}
	return nil
}
