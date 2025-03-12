package public

import (
	"blog_noe/database"
	"blog_noe/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := database.GetDB()

	var input struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.Email = input.Email
	user.Username = input.Username

	if err := db.Create(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Can't create user"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
