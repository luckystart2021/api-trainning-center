package usermodel

import (
	"api-trainning-center/models"
	"database/sql"
	"errors"
	"log"
)

// CreateUserByRequest executes subscribe to updates from an email address
func CreateUserByRequest(req models.AccountRequest, db *sql.DB) error {
	query := `INSERT INTO "user" (username, password, email, role) VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(query, req.UserName, req.PassWord, req.Email, req.Role)
	if err != nil {
		log.Println("Insert DB err", err)
		return err
	}
	return nil
}

func CheckUserLogin(req models.AccountRequest, db *sql.DB) (models.User, error) {
	user := models.User{}
	query := `
	SELECT 
		username, password, role
	FROM 
		"user" u 
	WHERE 
		u.username = $1;`
	row := db.QueryRow(query, req.UserName)
	err := row.Scan(&user.UserName, &user.PassWord, &user.Role)
	if err != nil {
		return user, errors.New("Login failed, please try again")
	}
	return user, nil
}
