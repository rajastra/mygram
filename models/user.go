package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null;type:varchar(100)"`
	Password string `gorm:"not null"`
	Age int `gorm:"not null;default:18;check:age>8"`
}