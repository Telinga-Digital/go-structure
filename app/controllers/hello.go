package controllers

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.PureJSON(200, gin.H{
		"status":  "Ok",
		"Message": "API Works!",
	})
}
