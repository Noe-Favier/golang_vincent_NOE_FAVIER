package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement" binding:"omitempty"`

	Content string `json:"content" binding:"required"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID" binding:"omitempty"`
	User    User   `json:"user" gorm:"foreignKey:UserID" binding:"omitempty"`

	PostID uint `json:"post_id" binding:"omitempty"`
	UserID uint `json:"user_id" binding:"omitempty"`
}
