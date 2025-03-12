package private

import (
	"blog_noe/database"
	"blog_noe/models"

	"github.com/gin-gonic/gin"
)

func UserPostsByMailHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	email := c.Query("email")

	var posts []models.Post

	if user.Email == email {
		//if ourselves return the user from the context
		c.JSON(200, user.Posts)
		return
	} else {

		db := database.GetDB()

		var author models.User
		if err := db.Preload("Posts").Where("email = ?", email).First(&author).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		posts = author.Posts
	}

	c.JSON(200, posts)
}

func UploadPostHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	db := database.GetDB()

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input.User = user
	input.UserID = user.ID

	if err := db.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, input)
}
