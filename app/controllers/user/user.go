package user

import "github.com/gin-gonic/gin"

func Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
		"data":   "",
	})
}
