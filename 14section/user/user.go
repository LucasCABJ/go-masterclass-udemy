package user

import (
	"errors"
	"strings"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
)

func CheckUsername(username string) bool {
	if len(username) < 6 || strings.Contains(username, "admin") {
		return false
	}
	return true
}

func Login(username, password string) (bool, error) {
	if !CheckUsername(username) {
		return false, ErrInvalidUsername
	}
	if strings.Compare(password, "1234") != 0 {
		return false, ErrInvalidPassword
	}
	return true, nil
}
