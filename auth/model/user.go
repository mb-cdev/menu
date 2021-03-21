package model

import (
	"crypto/sha512"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Login     string `gorm:"unique"`
	Password  []byte
	Address   []*Address
}

func (u *User) SetPassword(p string) {
	s := sha512.New512_256()

	s.Write([]byte(p))
	u.Password = s.Sum(nil)
}
