package model

import (
	"crypto/sha512"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname     string `validator:"notEmptyString"`
	Lastname      string `validator:"notEmptyString"`
	Login         string `gorm:"unique" validator:"login"`
	Email         string `validator:"email"`
	Password      []byte
	IsBackendUser bool
}

func NewUserFactory(firstname string, lastname string, login string, password string, email string, is_backend_user bool) *User {
	u := User{
		Firstname:     firstname,
		Lastname:      lastname,
		Login:         login,
		Email:         email,
		IsBackendUser: is_backend_user,
	}

	u.SetPassword(password)
	return &u
}

func (u *User) SetPassword(p string) {
	s := sha512.New()

	s.Write([]byte(p))
	u.Password = s.Sum(nil)
}
