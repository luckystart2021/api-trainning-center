package admin

import (
	"api-trainning-center/validate"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
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

type MessageResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Reponse struct {
	Status bool `json:"status"`
}

type LoginReponse struct {
	Success  bool   `json:"success"`
	Token    string `json:"token"`
	InfoUser User   `json:"infoUser"`
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
func UpdateAccountByRequest(userName, newPassWord string, db *sql.DB) error {
	query := `
	UPDATE "user" SET
		password=$1
	WHERE 
		username = $2;`
	_, err := db.Exec(query, newPassWord, userName)
	if err != nil {
		log.Println("Update DB err", err)
		return err
	}
	return nil
}

func RetrieveAccountByUserName(userName string, db *sql.DB) (User, error) {
	user := User{}
	query := `
	SELECT 
		username, "password", email, "role", sex, dateofbirth, phone, fullname, created_at, is_delete
	FROM 
		"user" u 
	WHERE 
		u.username = $1;`
	row := db.QueryRow(query, userName)

	err := row.Scan(&user.UserName, &user.PassWord, &user.Email, &user.Role, &user.Sex, &user.DateOfBirth, &user.Phone, &user.FullName, &user.CreatedAt, &user.IsDelete)
	if err != nil {
		log.Fatalln("RetrieveAccountByUserName scan error", err)
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
		"user"`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln("[RetrieveAccounts] query error ", err)
		return users, err
	}

	for rows.Next() {
		var err error
		var username, password, email, role, sex, dateofbirth, phone, fullname string
		var created_at time.Time
		var is_delete bool
		err = rows.Scan(&username, &password, &email, &role, &sex, &dateofbirth, &phone, &fullname, &created_at, &is_delete)
		if err != nil {
			log.Fatalln("[RetrieveAccounts] Scan error ", err)
			return users, err
		}
		user := User{
			UserName:    username,
			PassWord:    password,
			Role:        role,
			Email:       email,
			Sex:         sex,
			DateOfBirth: dateofbirth,
			Phone:       phone,
			FullName:    fullname,
			IsDelete:    is_delete,
		}
		user.CreatedAt, err = TimeIn(created_at, "Local")
		if err != nil {
			log.Println("<time unknown>")
		}
		users = append(users, user)
	}
	return users, nil
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
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
