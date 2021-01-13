package admin

import (
	"api-trainning-center/validate"
	"database/sql"
	"errors"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string
	PassWord string
	Role     string
}

type AccountRequest struct {
	Email       string `json:"email"`
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	Role        string `json:"role"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"dateofbirth"`
	Phone       string `json:"phone"`
	FullName    string `json:"fullname"`
}

type Reponse struct {
	Status bool `json:"status"`
}

type LoginReponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	UserID  string `json:"userid"`
}

type ChangeAccountRequest struct {
	UserName    string `json:"username"`
	OldPassWord string `json:"oldpassword"`
	NewPassWord string `json:"newpassword"`
}

func (acc AccountRequest) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if acc.UserName == "" {
			return errors.New("Bạn chưa nhập tên đăng nhập")
		} else {
			if len(acc.FullName) > 50 {
				return errors.New("Tên đăng nhập quá dài")
			}
		}
		if acc.PassWord == "" {
			return errors.New("Bạn chưa nhập mật khẩu")
		} else {
			if len(acc.PassWord) > 64 {
				return errors.New("Mật khẩu quá dài")
			}
		}
		return nil
	default: // this is for creating a user, where all fields are required
		if acc.UserName == "" {
			return errors.New("Bạn chưa nhập tên đăng nhập")
		} else {
			if len(acc.UserName) > 50 {
				return errors.New("Tên đăng nhập không được quá dài")
			}
		}

		if acc.PassWord == "" {
			return errors.New("Bạn chưa nhập mật khẩu")
		} else {
			if err := validate.VerifyPassword(acc.PassWord); err != nil {
				return err
			}
		}

		if acc.Role == "" {
			return errors.New("Bạn chưa nhập quyền")
		} else {
			if err := validate.IsValidRole(acc.Role); err != nil {
				return err
			}
		}

		if acc.Sex == "" {
			return errors.New("Bạn chưa nhập giới tính")
		} else {
			if len(acc.Sex) > 20 {
				return errors.New("Giới tính không được quá dài")
			}
		}

		if acc.DateOfBirth == "" {
			return errors.New("Bạn chưa nhập ngày tháng năm sinh")
		} else {
			if len(acc.DateOfBirth) > 10 || !validate.CheckDate(acc.DateOfBirth) {
				return errors.New("Ngày tháng năm sinh không đúng định dạng")
			}
		}

		if acc.Phone == "" {
			return errors.New("Bạn chưa nhập số điện thoại")
		} else {
			if len(acc.Phone) > 15 || !validate.CheckNumber(acc.Phone) {
				return errors.New("Số điện thoại không đúng")
			}
		}

		if acc.FullName == "" {
			return errors.New("Bạn chưa nhập họ và tên")
		} else {
			if len(acc.FullName) > 100 {
				return errors.New("Họ và tên quá dài")
			}
		}

		if acc.Email != "" {
			if err := checkmail.ValidateFormat(acc.Email); err != nil {
				return errors.New("Email không đúng định dạng")
			}
		}
	}

	return nil
}

func (accChange ChangeAccountRequest) ChangeAccountValidate() error {
	if accChange.UserName == "" {
		return errors.New("Bạn chưa nhập tên đăng nhập")
	}
	if accChange.OldPassWord == "" {
		return errors.New("Bạn chưa nhập mật khẩu cũ")
	}
	if accChange.OldPassWord == accChange.NewPassWord {
		return errors.New("Mật khẩu mới không được trùng với mật khẩu cũ")
	}
	if accChange.NewPassWord == "" {
		return errors.New("Bạn chưa nhập mật khẩu mới")
	} else {
		if err := validate.VerifyPassword(accChange.NewPassWord); err != nil {
			return err
		}
	}
	return nil
}

// CreateUserByRequest executes subscribe to updates from an email address
func CreateUserByRequest(req AccountRequest, db *sql.DB) error {
	query := `
	INSERT INTO "user" 
		(username, password, email, role, sex, dateofbirth, phone, fullname) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := db.Exec(query, req.UserName, req.PassWord, req.Email, req.Role, req.Sex, req.DateOfBirth, req.Phone, req.FullName)
	if err != nil {
		log.Println("Insert DB err", err)
		return err
	}
	return nil
}

// CreateUserByRequest executes subscribe to updates from an email address
func UpdateAccountByRequest(req ChangeAccountRequest, db *sql.DB) error {
	query := `
	UPDATE "user" SET
		username = $1, password=$2
	WHERE 
		username = $1;`
	_, err := db.Exec(query, req.UserName, req.NewPassWord)
	if err != nil {
		log.Println("Update DB err", err)
		return err
	}
	return nil
}

func CheckUserLogin(userName string, db *sql.DB) (User, error) {
	user := User{}
	query := `
	SELECT 
		username, password, role
	FROM 
		"user" u 
	WHERE 
		u.username = $1;`
	row := db.QueryRow(query, userName)
	err := row.Scan(&user.UserName, &user.PassWord, &user.Role)
	if err != nil {
		log.Println("CheckUserLogin error", err)
		return user, errors.New("Tên đăng nhập hoặc mật khẩu không đúng")
	}
	return user, nil
}

// HashPassword hashes password from user input
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 is the cost for hashing the password.
	if err != nil {
		log.Println("HashPassword error", err)
		return nil, errors.New("Đăng nhập thất bại")
	}
	return bytes, err
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("CheckPasswordHash error", err)
		return errors.New("Đăng nhập thất bại")
	}
	return nil
}
