package main

import (
	"blog_noe/database"
	"blog_noe/routers"
)

func main() {
	database.SetupDB()

	routers.New().SetupRouter().Run(":8080")
}
