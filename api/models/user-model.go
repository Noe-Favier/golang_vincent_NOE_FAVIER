package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement" binding:"omitempty"`

	Email    string `json:"email" gorm:"not null;uniqueIndex" binding:"required"`
	Username string `json:"username" gorm:"not null;uniqueIndex" binding:"required"`

	Follows []User `json:"-" gorm:"many2many:user_follows;" binding:"omitempty"`

	Posts    []Post    `json:"-" binding:"omitempty"`
	Comments []Comment `json:"-" binding:"omitempty"`
	Likes    []Post    `json:"-" binding:"omitempty"`
}
