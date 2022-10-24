package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Caption   string 
	Photo_url string `gorm:"not null"`
	UserRefer uint
}