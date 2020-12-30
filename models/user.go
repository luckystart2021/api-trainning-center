package models

import (
	"errors"

	"github.com/badoux/checkmail"
)

type AccountRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
}

type AccountReponse struct {
	Status bool `json:"status"`
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

func (acc AccountRequest) Validate() error {
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

	return nil
}
