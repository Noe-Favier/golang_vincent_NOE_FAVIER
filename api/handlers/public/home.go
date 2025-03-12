package public

import (
	"blog_noe/database"
	"blog_noe/models"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	db := database.GetDB()

	var posts []models.Post
	if err := db.Order("created_at desc").Limit(5).Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, posts)
}
