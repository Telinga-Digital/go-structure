package main

import (
	"github.com/Telinga-Digital/go-structure/app/models"
	"github.com/Telinga-Digital/go-structure/config"
	routes "github.com/Telinga-Digital/go-structure/routes/v1"
	"github.com/gin-gonic/gin"
)

var (
	Port = config.GetAppConfig().Port
)

func main() {
	router := gin.Default()

	models.MakeConnection()

	routes.GetAllRoutes(router)

	router.Run(":" + Port)
}
