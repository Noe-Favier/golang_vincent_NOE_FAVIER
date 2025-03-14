package middlewares

import (
	"blog_noe/database"
	"blog_noe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required (email)"})
			c.Abort()
			return
		}

		var user models.User
		if err := db.
			Preload("Likes").
			Preload("Posts").
			Preload("Posts").
			Preload("Posts.User").
			Preload("Posts.Likes").
			Preload("Posts.Comments").
			Preload("Posts.Comments.User").
			Preload("Comments").
			Preload("Comments.User").
			Where("email = ?", authorizationHeader).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
