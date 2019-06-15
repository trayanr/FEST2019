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
	ID       int
	Username string `json:"username"`
	Email    string
	Password string `json:"password"`
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
	if u.Email == "" {
		errs = append(errs, errInvalidEmail)
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
