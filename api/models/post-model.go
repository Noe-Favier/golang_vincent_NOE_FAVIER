package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	Title   string `json:"title"`
	Content string `json:"content"`

	User User `json:"user" gorm:"foreignKey:UserID"`

	Likes    []User    `json:"likes" gorm:"many2many:post_likes;"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`

	UserID uint `json:"user_id"`
}
