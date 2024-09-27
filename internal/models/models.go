package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Snippet struct {
	gorm.Model
	Content   string
	UserID    uint
	User      User
	ViewCount uint
}
