package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement" binding:"omitempty"`

	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`

	User   User `json:"user" gorm:"foreignKey:UserID" binding:"omitempty"` //omit empty bc user is found in context
	UserID uint `json:"user_id" binding:"omitempty"`

	Likes    []User    `json:"likes" gorm:"many2many:post_likes;" binding:"omitempty"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID" binding:"omitempty"`
}
