package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const (
	ADMIN   string = "ADMIN"
	TEACHER string = "TEACHER"
	EDITOR  string = "EDITOR"
)

func CheckDate(date string) bool {
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	return re.MatchString(date)
}

func CheckNumber(number string) bool {
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(number)
}

func VerifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 64
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("Mật khẩu phải có ít nhất 1 ký tự thường")
	}
	if !uppercasePresent {
		appendError("Mật khẩu phải có ít nhất 1 ký tự in hoa")
	}
	if !numberPresent {
		appendError("Mật khẩu phải có ít nhất 1 ký tự số")
	}
	if !specialCharPresent {
		appendError("Mật khẩu phải có ít nhất 1 ký tự đặc biệt")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("Độ dài mật khẩu phải dài từ %d đến %d ký tự", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

func IsValidRole(role string) error {
	switch role {

	case ADMIN:
		return nil

	case TEACHER:
		return nil

	case EDITOR:
		return nil
	}
	return errors.New("Phân quyền không tồn tại")
}
