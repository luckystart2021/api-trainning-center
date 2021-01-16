package account

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type User struct {
	UserName    string    `json:"username"`
	PassWord    string    `json:"password"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	Sex         string    `json:"sex"`
	Phone       string    `json:"phone"`
	FullName    string    `json:"fullname"`
	CreatedAt   time.Time `json:"created_at"`
	IsDelete    bool      `json:"is_delete"`
}

func RetrieveAccountByUserName(userName string, db *sql.DB) (User, error) {
	user := User{}
	query := `
	SELECT 
		username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, is_delete
	FROM 
		"user" u 
	WHERE 
		u.username = $1 AND is_delete = $2;`
	row := db.QueryRow(query, userName, false)

	var email sql.NullString

	err := row.Scan(&user.UserName, &user.PassWord, &email, &user.Role, &user.Sex, &user.DateOfBirth, &user.Phone, &user.FullName, &user.CreatedAt, &user.IsDelete)
	if email.Valid {
		user.Email = email.String
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("RetrieveAccountByUserName scan error  %v", err)
		return user, errors.New("Tên đăng nhập không đúng")
	}
	return user, nil
}

func RetrieveAccounts(db *sql.DB) ([]User, error) {
	users := []User{}
	query := `
	SELECT 
		username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, is_delete
	FROM 
		"user"
	WHERE 
		is_delete = $1;`
	rows, err := db.Query(query, false)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[RetrieveAccounts] query error  %v", err)
		return users, err
	}

	for rows.Next() {
		var err error
		var email sql.NullString
		var username, password, role, sex, dateofbirth, phone, fullname string
		var created_at time.Time
		var is_delete bool
		err = rows.Scan(&username, &password, &email, &role, &sex, &dateofbirth, &phone, &fullname, &created_at, &is_delete)
		if err != nil {
			logrus.WithFields(logrus.Fields{}).Errorf("[RetrieveAccounts] Scan error  %v", err)
			return users, err
		}
		user := User{
			UserName:    username,
			PassWord:    password,
			Role:        role,
			Sex:         sex,
			CreatedAt:   created_at,
			DateOfBirth: dateofbirth,
			Phone:       phone,
			FullName:    fullname,
			IsDelete:    is_delete,
		}

		if email.Valid {
			user.Email = email.String
		}
		users = append(users, user)
	}
	return users, nil
}
