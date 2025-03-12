package private

import (
	"blog_noe/database"
	"blog_noe/models"
	"net/http"
	"strconv"

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
		if err := db.
			Preload("Posts").
			Preload("Posts.User").
			Preload("Posts.Likes").
			Preload("Posts.Comments").
			Preload("Posts.Comments.User").
			Where("email = ?", email).
			First(&author).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
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

func AddCommentHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	post_id := c.Param("post_id")
	db := database.GetDB()

	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input.User = user
	input.UserID = user.ID
	postID, err := strconv.Atoi(post_id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid post ID"})
		return
	}
	input.PostID = uint(postID)

	if err := db.Create(&input).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, "Comment added successfully")
}

func LikePostHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := database.GetDB()
	id := c.Param("post_id")

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	post.Likes = append(post.Likes, user)

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

func UnlikePostHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := database.GetDB()
	id := c.Param("post_id")

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, u := range post.Likes {
		if u.ID == user.ID {
			post.Likes = append(post.Likes[:i], post.Likes[i+1:]...)
			break
		}
	}

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}

func DeletePostHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	db := database.GetDB()
	id := c.Param("post_id")

	var post models.Post
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if post.UserID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not the author of this post"})
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
