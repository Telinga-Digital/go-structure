package hello

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"status":  "Ok",
		"Message": "API Works!",
	})
}
