package model

import (
	"crypto/sha512"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname     string
	Lastname      string
	Login         string `gorm:"unique" validator:"login"`
	Password      []byte
	IsBackendUser bool
}

func NewUserFactory(firstname string, lastname string, login string, password string, is_backend_user bool) *User {
	u := User{
		Firstname:     firstname,
		Lastname:      lastname,
		Login:         login,
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
