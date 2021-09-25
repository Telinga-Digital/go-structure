package routes

import (
	"github.com/endrureza/go-structure/app/controllers"
	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine) {
	router.GET("/", controllers.Hello)
}
