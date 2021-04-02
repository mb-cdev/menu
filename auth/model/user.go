package model

import (
	"crypto/sha512"
	"errors"
	"menu/common/database"
	"menu/common/validator"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname     string `validator:"notEmptyString"`
	Lastname      string `validator:"notEmptyString"`
	Login         string `gorm:"unique" validator:"login"`
	Email         string `validator:"email" gorm:"unique"`
	Password      []byte
	IsBackendUser bool
}

func NewUserFactory(firstname string, lastname string, login string, password string, email string, is_backend_user bool) (*User, error) {
	u := User{
		Firstname:     firstname,
		Lastname:      lastname,
		Login:         login,
		Email:         email,
		IsBackendUser: is_backend_user,
	}
	/**
	 * Validate password before hash
	 **/
	if v := validator.IsPasswordValid([]byte(password)); !v {
		return nil, errors.New("Password is invalid")
	}
	/**
	 * Check if user exits in database
	 **/

	var c int64
	database.DB.Model(&User{}).
		Where(database.DB.Where("email = ?", email).Or("login = ?", login)).
		Count(&c)
	if c > 0 {
		return nil, errors.New("User exists")
	}

	u.SetPassword(password)
	return &u, nil
}

func (u *User) SetPassword(p string) {
	s := sha512.New()

	s.Write([]byte(p))
	u.Password = s.Sum(nil)
}
