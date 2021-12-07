package middlewares

import (
	"github.com/Telinga-Digital/go-structure/app/models"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		apiKey := c.Query("api_key")

		if apiKey == "" {
			c.AbortWithStatusJSON(422, gin.H{
				"message": "API Key is required!",
			})
			return
		}

		var user models.User

		result := models.DB.Where("api_key = ?", apiKey).First(&user)

		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "API Key is invalid!",
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}

	return gin.HandlerFunc(fn)
}
