package account

import (
	"api-trainning-center/validate"
	"database/sql"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AccountRequest struct {
	Email       string `json:"email"`
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	Role        string `json:"role"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"dateofbirth"`
	Phone       string `json:"phone"`
	FullName    string `json:"fullname"`
	Address     string `json:"address"`
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
			if len(acc.UserName) > 50 {
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

		if action != "update" {
			if acc.PassWord == "" {
				return errors.New("Bạn chưa nhập mật khẩu")
			} else {
				if err := validate.VerifyPassword(acc.PassWord); err != nil {
					return err
				}
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

		if acc.Address == "" {
			return errors.New("Bạn chưa nhập địa chỉ")
		} else {
			if len(acc.Address) > 250 {
				return errors.New("Địa chỉ quá dài")
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
	INSERT INTO "users" 
		(username, password, email, role, sex, dateofbirth, phone, fullname, address) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	_, err := db.Exec(query, req.UserName, req.PassWord, req.Email, req.Role, req.Sex, req.DateOfBirth, req.Phone, req.FullName, req.Address)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CreateUserByRequest]Insert DB err  %v", err)
		return err
	}
	return nil
}

// HashPassword hashes password from user input
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 is the cost for hashing the password.
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[HashPassword] error  %v", err)
		return nil, errors.New("Đăng nhập thất bại")
	}
	return bytes, err
}

// CheckPasswordHash checks password hash and password from user input if they match
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[CheckPasswordHash] error  %v", err)
		return errors.New("Đăng nhập thất bại")
	}
	return nil
}
