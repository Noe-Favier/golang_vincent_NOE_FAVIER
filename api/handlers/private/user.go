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
