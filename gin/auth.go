package user

import (
	"errors"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

func CheckUserIsExist(username string) bool {
	_, isExist := UserData[username]
	return isExist
}

func CheckPassword(pwd1 string, pwd2 string) error {
	if pwd1 == pwd2 {
		return nil
	} else {
		return errors.New("password mismatch")
	}
}

func Auth(account string, password string) error {
	if isExist := CheckUserIsExist(account); isExist {
		return CheckPassword(UserData[username], password)
	} else {
		return errors.New("user not exist")
	}
}
