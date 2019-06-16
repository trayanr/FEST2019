package models

import (
	"errors"

	"github.com/kachamaka/argon2custom"
)

var (
	errInvalidName     = errors.New("invalid username")
	errInvalidPassword = errors.New("invalid password")
	errInvalidEmail    = errors.New("invalid email")
)

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	OAuthCode   string
	Level       int  `json:"level"`
	Points      int  `json:"points"`
	LastChecked uint `json:"lastChecked"`
}

func (u User) Validate() (errs []error, ok bool) {
	ok = true
	if u.Username == "" {
		errs = append(errs, errInvalidName)
		ok = false
	}
	if u.Password == "" {
		errs = append(errs, errInvalidPassword)
		ok = false
	}
	return
}

func (u *User) HashPassword() error {
	hash, err := argon2custom.GenerateFromPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}
