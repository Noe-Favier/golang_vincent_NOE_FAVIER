package database

import (
	"blog_noe/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() {
	//TODO: add to .env
	host := "localhost"
	user := "gorm"
	password := "gorm"
	dbname := "blog_noe"
	port := "5432"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
}
