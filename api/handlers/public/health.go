package public

import "github.com/gin-gonic/gin"

//	@Summary		Health check endpoint
//	@Description	Get API health status
//	@Tags			public
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/public/health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}
