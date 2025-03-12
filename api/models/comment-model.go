package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	Content string `json:"content"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`

	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}
