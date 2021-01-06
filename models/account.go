package models

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type AccessDetails struct {
	Role     string
	UserName string
}

type User struct {
	UserName string
	PassWord string
	Role     string
}

type AccountRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
}

type AccountReponse struct {
	Status bool `json:"status"`
}

type LoginReponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

const (
	ADMIN   string = "ADMIN"
	TEACHER string = "TEACHER"
	EDITOR  string = "EDITOR"
)

func (acc AccountRequest) IsValid() (bool, error) {
	switch acc.Role {

	case ADMIN:
		return true, nil

	case TEACHER:
		return true, nil

	case EDITOR:
		return true, nil
	}

	return false, errors.New("Role does not exist")
}

func (acc AccountRequest) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if acc.UserName == "" {
			return errors.New("UserName is required")
		}
		if acc.PassWord == "" {
			return errors.New("Password is required")
		}
		return nil
	default: // this is for creating a user, where all fields are required
		if acc.UserName == "" {
			return errors.New("Required UserName")
		}

		if acc.PassWord == "" {
			return errors.New("Required Password")
		}

		if acc.Role == "" {
			return errors.New("Required Role")
		}

		if acc.Email != "" {
			if err := checkmail.ValidateFormat(acc.Email); err != nil {
				return errors.New("Invalid Email")
			}
		}
	}

	return nil
}

// CreateUserByRequest executes subscribe to updates from an email address
func CreateUserByRequest(req AccountRequest, db *sql.DB) error {
	query := `INSERT INTO "user" (username, password, email, role) VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(query, req.UserName, req.PassWord, req.Email, req.Role)
	if err != nil {
		log.Println("Insert DB err", err)
		return err
	}
	return nil
}

func CheckUserLogin(req AccountRequest, db *sql.DB) (User, error) {
	user := User{}
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

// HashPassword hashes password from user input
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 is the cost for hashing the password.
	if err != nil {
		return nil, errors.New("hashes password error")
	}
	return bytes, err
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}
