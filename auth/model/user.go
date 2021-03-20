package model

type User struct {
	ID        uint `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	Login     string
	password  string
}
