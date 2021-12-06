package v1

import (
	"github.com/Telinga-Digital/go-structure/app/controllers/hello"
	"github.com/Telinga-Digital/go-structure/app/controllers/user"
	"github.com/gin-gonic/gin"
)

func GetAllRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/", hello.Index)

		v1.GET("/user", user.Show)
	}
}
