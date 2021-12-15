package user

import (
	"encoding/json"

	ExampleJob "github.com/Telinga-Digital/go-structure/jobs/example_job"
	"github.com/Telinga-Digital/go-structure/services/redis"
	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
		"data":   c.MustGet("user"),
	})
}

func Broadcast(c *gin.Context) {
	user, _ := json.Marshal(c.MustGet("user"))

	redis.Publish(ExampleJob.Channel, user)

	c.JSON(200, gin.H{
		"status": "OK",
		"data":   "Broadcast User Success!",
	})
}
