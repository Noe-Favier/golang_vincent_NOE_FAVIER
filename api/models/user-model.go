package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	Email    string `json:"-" gorm:"not null;uniqueIndex"`
	Username string `json:"username" gorm:"not null;uniqueIndex"`

	Follows []User `json:"-" gorm:"many2many:user_follows;"`

	Posts    []Post    `json:"-"`
	Comments []Comment `json:"-"`
	Likes    []Post    `json:"-"`
}
