package main

import (
	"github.com/endrureza/go-structure/config"
	"github.com/endrureza/go-structure/routes"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	Port   = config.GetAppConfig().Port
)

func main() {
	router := gin.Default()

	routes.GetRoutes(router)

	if Port == "" {
		endless.ListenAndServe(":8080", router)
	} else {
		endless.ListenAndServe(":"+Port, router)
	}
}
