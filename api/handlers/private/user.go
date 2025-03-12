package private

import (
	"blog_noe/database"
	"blog_noe/models"

	"github.com/gin-gonic/gin"
)

func GetUserByID(id uint, c *gin.Context) (models.User, error) {
	user := c.MustGet("user").(models.User)

	if user.ID == id {
		//if ourselves return the user from the context
		return user, nil
	}

	db := database.GetDB()

	var user_find models.User
	if err := db.First(&user_find, id).Error; err != nil {
		return user, err
	}

	return user_find, nil
}

func FollowUserByIDHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	user_id := c.Param("user_id")

	db := database.GetDB()

	var user_find models.User
	if err := db.First(&user_find, user_id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := db.Model(&user).Association("Following").Append(&user_find); err != nil {
		c.JSON(500, gin.H{"error": "Could not follow user"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}

func UnfollowUserByIDHandler(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	user_id := c.Param("user_id")

	db := database.GetDB()

	var user_find models.User

	if err := db.First(&user_find, user_id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := db.Model(&user).Association("Following").Delete(&user_find); err != nil {
		c.JSON(500, gin.H{"error": "Could not unfollow user"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
