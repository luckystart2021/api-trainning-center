package database

import (
	"api-trainning-center/models"
	"log"
)

// CreateSubscribeFriendByRequestorAndTarget executes subscribe to updates from an email address
func (db Database) CreateUserByRequest(req models.AccountRequest) error {
	query := `INSERT INTO "user" (username, password, email, role) VALUES ($1, $2, $3, $4);`
	_, err := db.Conn.Exec(query, req.UserName, req.PassWord, req.Email, req.Role)
	if err != nil {
		log.Println("Insert DB err", err)
		return err
	}
	return nil
}
